package four

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

var (
	regexHairColor = regexp.MustCompile("^#[0-9a-f]{6}$")
	regexEyeColor  = regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
)

func (p *Passport) Validate() bool {

	// four digits; at least 1920 and at most 2002.
	if p.BirthYear > 2002 || p.BirthYear < 1920 {
		return false
	}

	// four digits; at least 2010 and at most 2020.
	if p.IssueYear > 2020 || p.IssueYear < 2010 {
		return false
	}

	// four digits; at least 2020 and at most 2030.
	if p.ExpirationYear > 2030 || p.ExpirationYear < 2020 {
		return false
	}

	// Height
	// a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	{
		if p.Height == "" {
			return false
		}
		hightSufix := p.Height[len(p.Height)-2:]
		if hightSufix != "cm" && hightSufix != "in" {
			return false
		}
		number := p.Height[:len(p.Height)-2]
		i, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			i = -1
		}
		if hightSufix == "cm" && (i < 150 || i > 193) {
			return false
		} else if hightSufix == "in" && (i < 59 || i > 76) {
			return false
		}
	}

	// Hair Color
	// a # followed by exactly six characters 0-9 or a-f
	if !regexHairColor.MatchString(p.HairColor) {
		return false
	}

	// exactly one of: amb blu brn gry grn hzl oth.
	if !regexEyeColor.MatchString(p.EyeColor) {
		return false
	}

	// Passport ID
	//a nine-digit number, including leading zeroes.
	{
		if len(p.PassportID) != 9 {
			return false
		}

		_, err := strconv.ParseInt(p.PassportID, 10, 64)
		if err != nil {
			return false
		}

	}

	// if p.CountryID == "" {
	// 	return false
	// }

	return true
}

func BatchValidatePassports(s string) int {
	passports := ParseOutIndividualPassports(s)

	count := 0

	for _, p := range passports {
		if p.Validate() {
			count++
		}
	}

	return count
}

func ParseOutIndividualPassports(s string) []Passport {

	passports := make([]Passport, 0)

	lines := strings.Split(s, "\n")

	b := strings.Builder{}

	for _, l := range lines {
		clean := strings.TrimSpace(l)
		if clean == "" {
			p := ParsePassport(b.String())
			b.Reset()
			passports = append(passports, p)
			continue
		}

		b.WriteString(clean)
		b.WriteString(" ")
	}

	p := ParsePassport(b.String())
	passports = append(passports, p)

	return passports
}

func ParsePassport(s string) Passport {

	s = strings.TrimSpace(s)
	p := Passport{}
	vals := strings.Split(s, " ")

	for _, field := range vals {
		split := strings.Split(field, ":")
		f := split[0]
		v := split[1]
		switch f {
		case "byr":
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				i = -1
			}
			p.BirthYear = int(i)
		case "iyr":
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				i = -1
			}
			p.IssueYear = int(i)
		case "eyr":
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				i = -1
			}
			p.ExpirationYear = int(i)
		case "hgt":
			p.Height = v
		case "hcl":
			p.HairColor = v
		case "ecl":
			p.EyeColor = v
		case "pid":
			p.PassportID = v
		case "cid":
			p.CountryID = v
		default:
			fmt.Printf("Unknown Field: %s:%s\n", f, v)
		}
	}

	return p
}
