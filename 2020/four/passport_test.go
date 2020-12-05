package four

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePassport(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want Passport
	}{
		{
			"Example -1",
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
		},
		{
			"Example - 2",
			"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
			Passport{
				BirthYear:      1929,
				IssueYear:      2013,
				ExpirationYear: 2023,
				// Height:         "", Missing height
				HairColor:  "#cfa07d",
				EyeColor:   "amb",
				PassportID: "028048884",
				CountryID:  "350",
			},
		},
		{
			"Example - 3",
			"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
			Passport{
				BirthYear:      1931,
				IssueYear:      2013,
				ExpirationYear: 2024,
				Height:         "179cm",
				HairColor:      "#ae17e1",
				EyeColor:       "brn",
				PassportID:     "760753108",
				// CountryID:      "147", Missing CID
			},
		},
		{
			"Example - 4",
			"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
			Passport{
				// BirthYear:      "1937", Missing bry
				IssueYear:      2011,
				ExpirationYear: 2025,
				Height:         "59in",
				HairColor:      "#cfa07d",
				EyeColor:       "brn",
				PassportID:     "166559648",
				// CountryID:      "147",  Missing cid
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParsePassport(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseOutIndividualPassports(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			"example",
			example,
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseOutIndividualPassports(tt.s)
			assert.Len(t, got, tt.want)

		})
	}
}

func TestPassport_Validate(t *testing.T) {

	tests := []struct {
		name string
		p    Passport
		want bool
	}{
		{
			"Example -1",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			true,
		},
		{
			"Example - 2",
			Passport{
				BirthYear:      1929,
				IssueYear:      2013,
				ExpirationYear: 2023,
				// Height:         "", Missing height
				HairColor:  "#cfa07d",
				EyeColor:   "amb",
				PassportID: "028048884",
				CountryID:  "350",
			},
			false,
		},
		{
			"Example - 3",
			Passport{
				BirthYear:      1931,
				IssueYear:      2013,
				ExpirationYear: 2024,
				Height:         "179cm",
				HairColor:      "#ae17e1",
				EyeColor:       "brn",
				PassportID:     "760753108",
				// CountryID:      "147", Missing CID
			},
			true,
		},
		{
			"Example - 4",
			Passport{
				// BirthYear:      "1937", Missing bry
				IssueYear:      2011,
				ExpirationYear: 2025,
				Height:         "59in",
				HairColor:      "#cfa07d",
				EyeColor:       "brn",
				PassportID:     "166559648",
				// CountryID:      "147",  Missing cid
			},
			false,
		},
		{
			"Invalid BirthYear - high",
			Passport{
				BirthYear:      2003,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"Invalid BirthYear - low",
			Passport{
				BirthYear:      1919,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"Invalid IssueYear - low",
			Passport{
				BirthYear:      1937,
				IssueYear:      2009,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"Invalid IssueYear - high",
			Passport{
				BirthYear:      1937,
				IssueYear:      2021,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"Invalid ExpirationYear - low",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2019,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"Invalid ExpirationYear - high",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2031,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"Invalid Hight - Wrong sufix",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183ft",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"Invalid Hight - No sufix",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"Invalid Hight - Inch High",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183ft",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "77in",
			},
			false,
		},
		{
			"Invalid Hight - Inch Low",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183ft",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "58in",
			},
			false,
		},
		{
			"Invalid Hight - CM High",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183ft",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "194cm",
			},
			false,
		},
		{
			"Invalid Hight - CM Low",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183ft",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "149cm",
			},
			false,
		},
		{
			"Invalid Hair Color - no #",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "fffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"Invalid Hair Color - to many",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#ffffffffd",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"Invalid Hair Color - bad char",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#123abz",
				EyeColor:       "gry",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"EyeColor invalid - bad input",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "wat",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"EyeColor invalid - multiple input",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "ambblu",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"EyeColor invalid - no value",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "",
				PassportID:     "860033327",
				CountryID:      "147",
			},
			false,
		},
		{
			"Passport ID - bad length",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "86003332712",
				CountryID:      "147",
			},
			false,
		},
		{
			"Passport ID - leading 0 still count",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "000000001",
				CountryID:      "147",
			},
			true,
		},
		{
			"Passport ID - bad length",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "0123456789",
				CountryID:      "147",
			},
			false,
		},
		{
			"Passport ID - NaN",
			Passport{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2020,
				Height:         "183cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PassportID:     "A12345678",
				CountryID:      "147",
			},
			false,
		},
		// {
		// 	"Example -1",
		// 	Passport{
		// 		BirthYear:      1937,
		// 		IssueYear:      2017,
		// 		ExpirationYear: 2020,
		// 		Height:         "183cm",
		// 		HairColor:      "#fffffd",
		// 		EyeColor:       "gry",
		// 		PassportID:     "860033327",
		// 		CountryID:      "147",
		// 	},
		// 	true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.p.Validate()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBatchValidatePassports(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			"Example",
			example,
			2,
		},
		{
			"Solution",
			solution,
			167,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BatchValidatePassports(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
