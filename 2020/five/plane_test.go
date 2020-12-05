package five

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTicket(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			"example - 1",
			"FBFBBFFRLR",
			357,
		},
		{
			"example - 2",
			"BFFFBBFRRR",
			567,
		},
		{
			"example - 3",
			"FFFBBBFRRR",
			119,
		},
		{
			"example - 4",
			"BBFFBBFRLL",
			820,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseTicket(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHighestTicket(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			"example",
			`FBFBBFFRLR
			BFFFBBFRRR
			FFFBBBFRRR
			BBFFBBFRLL
			`,
			820,
		},
		{
			"solution",
			solution,
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HighestTicket(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFindEmptySeats(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			"solution",
			solution,
			554,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindEmptySeats(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
