package four

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validPasswordsInRange(t *testing.T) {
	min := 130254
	max := 678275

	c := ValidPasswordInRange(min, max)

	assert.Equal(t, 1419, c)
}

func Test_checkIsValid(t *testing.T) {
	testCases := []struct {
		numb  int
		valid bool
	}{
		{112345, true},
		{122456, true},
		{123356, true},
		{123446, true},
		{123455, true},
		{111345, false}, // more than 2 '1' in a row
		{213456, false}, // 1 < 2
		{123, false},    // Too short
		{123456, false}, //Needs xx numbe
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.valid, checkIsValid(tc.numb, checkIsSixDigits, checkRepeatingDigits, checkDoesNotDecrease))
		})
	}
}

func Test_checkRepeatingDigits(t *testing.T) {
	testCases := []struct {
		number int
		valid  bool
	}{

		{112345, true},
		{122456, true},
		{123356, true},
		{123446, true},
		{123455, true},

		{111155, true},

		{111345, false},
		{122256, false},
		{123336, false},
		{123444, false},
		{123555, false},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.valid, checkRepeatingDigits(tc.number))
		})
	}
}
func Test_digit(t *testing.T) {
	testCases := []struct {
		numb     int
		position int
		expected int
	}{
		{123, 1, 3},
		{123, 2, 2},
		{123, 3, 1},
		{123456789, 7, 3},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			o := digit(tc.numb, tc.position)
			assert.Equal(t, tc.expected, o)
		})
	}
}
