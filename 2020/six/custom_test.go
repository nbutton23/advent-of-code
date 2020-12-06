package six

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountGroupQuestions(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			"example 1",
			"abc",
			3,
		},
		{
			"example 2",
			`a
			b
			c`,
			0,
		},
		{
			"example 3",
			`ab
			ac`,
			1,
		},
		{
			"example 4",
			`a
			a
			a
			a`,
			1,
		},
		{
			"example 5",
			`b`,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountGroupQuestions(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCountAllGroupsQuestions(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want []int
	}{
		{
			"Example",
			example,
			[]int{3, 0, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountAllGroupsQuestions(tt.s)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestCount(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			"Example",
			example,
			6,
		},
		{
			"Solution 1",
			solution,
			//6427
			3125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Count(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
