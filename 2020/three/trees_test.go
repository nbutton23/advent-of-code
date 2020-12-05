package three

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_slop_correctX(t *testing.T) {

	tests := []struct {
		name    string
		xLenght int
		x       int
		want    int
	}{
		{
			"in bounds",
			11,
			3,
			3,
		},
		{
			"lenght past bounds",
			11,
			11,
			0,
		},
		{
			"two past bounds",
			11,
			12,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &slop{
				xLength: tt.xLenght,
			}
			got := s.correctX(tt.x)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_slop_isOffSlop(t *testing.T) {

	type args struct {
	}
	tests := []struct {
		name    string
		yLength int
		y       int
		want    bool
	}{
		{
			"On slop",
			11,
			5,
			false,
		},
		{
			"Off slop",
			11,
			11,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &slop{
				yLength: tt.yLength,
			}
			got := s.isOffSlop(tt.y)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parsePatternIntoSlop(t *testing.T) {

	tests := []struct {
		name    string
		s       string
		xLength int
		yLength int
	}{

		{
			"example",
			example,
			11,
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parsePatternIntoSlop(tt.s)
			assert.Equal(t, tt.xLength, got.xLength)
			assert.Equal(t, tt.yLength, got.yLength)
			assert.NotNil(t, got.pattern)
			assert.Len(t, got.pattern, 11)
			assert.Len(t, got.pattern[0], 11)
		})
	}
}

func Test_slop_getSpot(t *testing.T) {
	tests := []struct {
		name string
		p    string
		x    int
		y    int
		want string
	}{
		{
			"example -1",
			example,
			3,
			1,
			".",
		},
		{
			"example - 2",
			example,
			6,
			2,
			"#",
		},
		{
			"example - 3",
			example,
			9,
			3,
			".",
		},
		{
			"example - 4",
			example,
			12,
			4,
			"#",
		},
		{
			"example - 5",
			example,
			15,
			5,
			"#",
		},
		{
			"example - 6",
			example,
			18,
			6,
			".",
		},
		{
			"example - 7",
			example,
			21,
			7,
			"#",
		},
		{
			"example - 8",
			example,
			24,
			8,
			"#",
		},
		{
			"example - 9",
			example,
			27,
			9,
			"#",
		},
		{
			"example - 10",
			example,
			30,
			10,
			"#",
		},
		{
			"example off the map",
			example,
			33,
			11,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := parsePatternIntoSlop(tt.p)
			got := s.getSpot(tt.x, tt.y)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_slop_CountTreesInSlop(t *testing.T) {

	type args struct {
	}
	tests := []struct {
		name string
		p    string
		x    int
		y    int
		want int
	}{
		{
			"example",
			example,
			3,
			1,
			7,
		},
		{
			"solution 1",
			solution1,
			3,
			1,
			228,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := parsePatternIntoSlop(tt.p)

			got := s.CountTreesInSlop(tt.x, tt.y)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_slop_TreesInSlops(t *testing.T) {
	tests := []struct {
		name  string
		p     string
		slops []slopePair
		want  []int
	}{
		{
			"example",
			example,
			[]slopePair{
				{1, 1},
				{3, 1},
				{5, 1},
				{7, 1},
				{1, 2},
			},
			[]int{2, 7, 3, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := parsePatternIntoSlop(tt.p)

			got := s.CountTreesInSlops(tt.slops)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func Test_slop_ProductOfTreesInSlops(t *testing.T) {
	tests := []struct {
		name  string
		p     string
		slops []slopePair
		want  int
	}{
		{
			"example",
			example,
			[]slopePair{
				{1, 1},
				{3, 1},
				{5, 1},
				{7, 1},
				{1, 2},
			},
			336,
		},
		{
			"solution",
			solution1,
			[]slopePair{
				{1, 1},
				{3, 1},
				{5, 1},
				{7, 1},
				{1, 2},
			},
			6818112000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := parsePatternIntoSlop(tt.p)

			got := s.ProductOfTreesInSlops(tt.slops)
			assert.Equal(t, tt.want, got)
		})
	}
}
