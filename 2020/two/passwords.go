package two

import (
	"fmt"
	"regexp"
	"strconv"
)

var regexInt = regexp.MustCompile(`\d+`)
var regexStr = regexp.MustCompile("[a-zA-Z]+")

func parseString(in string) (string, int) {
	s := regexStr.FindString(in)
	c := regexStr.FindStringIndex(in)
	return s, c[1]
}

func parseOutInt(in string) (intFound, index int) {
	s := regexInt.FindString(in)
	c := regexInt.FindStringIndex(in)
	
	i, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		fmt.Println(err)
		return -1, -1

	}

	return int(i), c[1]
}

func parseInput(s string) (min, max int, char, password string) {
	if len(s) < 8 {
		return -1, -1, "", ""
	}
	//{min} - {max} {char}: {password}
	//1-3 b: cdefg

	min, c := parseOutInt(s)
	max, c2 := parseOutInt(s[c:])
	c = c + c2
	char, c3 := parseString(s[c:])
	c = c + c3
	pass, _ := parseString(s[c:])	

	return min, max, char, pass
}

func validatePasswordV1(password string, min, max int, char string) bool {

	count := 0

	for i := range password {
		if string(password[i]) == char {
			count++
		}
	}

	if count >= min && count <= max {
		return true
	}

	return false
}

func validatePasswordV2(password string, f, s int, char string) bool {

	if len(password) < s {
		fmt.Printf("index out of bounds for %s: %d", password, s)
		return false
	}

	p1 := string(password[f-1]) == char
	p2 := string(password[s-1]) == char

	// X xor Y <-> X != Y
	return p1 != p2
}

func HowManyValidePasswordsV1(list []string) int {
	count := 0
	for _, p := range list {
		min, max, char, password := parseInput(p)
		if char == "" {
			fmt.Printf("issue parseing %s\n", p)
			continue
		}

		if validatePasswordV1(password, min, max, char) {
			count++
		}
	}

	return count
}

func HowManyValidPasswordsV2(list []string) int {
	count := 0
	for _, p := range list {
		min, max, char, password := parseInput(p)
		if char == "" {
			fmt.Printf("issue parseing %s\n", p)
			continue
		}

		if validatePasswordV2(password, min, max, char) {
			count++
		}
	}

	return count
}
