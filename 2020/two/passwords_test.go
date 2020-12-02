package two

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name         string
		s            string
		wantMin      int
		wantMax      int
		wantChar     string
		wantPassword string
	}{
		{
			"example 1",
			"1-3 a: abcde",
			1,
			3,
			"a",
			"abcde",
		},
		{
			"example 2",
			"1-3 b: cdefg",
			1,
			3,
			"b",
			"cdefg",
		},
		{
			"example 3",
			"2-9 c: ccccccccc",
			2,
			9,
			"c",
			"ccccccccc",
		},
		{
			"case from solution",
			"11-12 n: frpknnndpntnncnnnnn",
			11,
			12,
			"n",
			"frpknnndpntnncnnnnn",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax, gotChar, gotPassword := parseInput(tt.s)
			assert.Equal(t, tt.wantMin, gotMin)
			assert.Equal(t, tt.wantMax, gotMax)
			assert.Equal(t, tt.wantChar, gotChar)
			assert.Equal(t, tt.wantPassword, gotPassword)

		})
	}
}

func Test_validatePasswordV1(t *testing.T) {

	tests := []struct {
		name     string
		password string
		min      int
		max      int
		char     string
		want     bool
	}{
		{
			"example 1",
			"abcde",
			1,
			3,
			"a",
			true,
		},
		{
			"example 2",
			"cdefg",
			1,
			3,
			"b",
			false,
		},
		{
			"example 3",
			"ccccccccc",
			2,
			9,
			"c",
			true,
		},
		{
			"case from solution",
			"frpknnndpntnncnnnnn",
			11,
			12,
			"n",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validatePasswordV1(tt.password, tt.min, tt.max, tt.char)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_validatePasswordV2(t *testing.T) {

	tests := []struct {
		name     string
		password string
		f        int
		s        int
		char     string
		want     bool
	}{
		{
			"example 1",
			"abcde",
			1,
			3,
			"a",
			true,
		},
		{
			"example 2",
			"cdefg",
			1,
			3,
			"b",
			false,
		},
		{
			"example 3",
			"ccccccccc",
			2,
			9,
			"c",
			false,
		},
		{
			"case from solution",
			"frpknnndpntnncnnnnn",
			11,
			12,
			"n",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validatePasswordV2(tt.password, tt.f, tt.s, tt.char)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHowManyValidePasswordsV1(t *testing.T) {
	tests := []struct {
		name string
		list []string
		want int
	}{
		{
			"example",
			[]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"},
			2,
		},
		{
			"solution 1",
			solutionInput,
			445,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HowManyValidePasswordsV1(tt.list)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHowManyValidPasswordsV2(t *testing.T) {

	tests := []struct {
		name string
		list []string
		want int
	}{
		{
			"example",
			[]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"},
			1,
		},
		{
			"solution 1",
			solutionInput,
			491,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HowManyValidPasswordsV2(tt.list)
			assert.Equal(t, tt.want, got)
		})
	}
}
