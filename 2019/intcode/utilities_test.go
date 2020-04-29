package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setValue(t *testing.T) {
	testCases := []struct {
		initProgram []int
		index       int
		value       int
		endLength   int
	}{
		{
			[]int{1, 1},
			1,
			3,
			2,
		},
		{
			initProgram: []int{1, 1},
			index:       2,
			value:       3,
			endLength:   4,
		},
		{
			initProgram: []int{1, 1, 1, 1},
			index:       5,
			value:       3,
			endLength:   8,
		},
		{
			initProgram: []int{1, 1},
			index:       5,
			value:       3,
			endLength:   6,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			p := &Process{
				program: tc.initProgram,
			}

			p.setValue(tc.index, tc.value)
			assert.Equal(t, tc.value, p.program[tc.index])
			assert.Len(t, p.program, tc.endLength)
		})
	}
}

func Test_getValue(t *testing.T) {
	testCases := []struct {
		initProgram []int
		index       int
		value       int
		endLength   int
	}{
		{
			initProgram: []int{1, 12},
			index:       1,
			value:       12,
			endLength:   2,
		},
		{
			initProgram: []int{1, 1},
			index:       2,
			value:       0,
			endLength:   4,
		},
		{
			initProgram: []int{1, 1, 1, 1},
			index:       5,
			value:       0,
			endLength:   8,
		},
		{
			initProgram: []int{1, 1},
			index:       5,
			value:       0,
			endLength:   6,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			p := &Process{
				program: tc.initProgram,
			}

			p.getValue(tc.index)
			assert.Equal(t, tc.value, p.program[tc.index])
			assert.Len(t, p.program, tc.endLength)
		})
	}
}
