package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_adjustRelativeBase(t *testing.T) {
	testCases := []struct {
		program     []int
		m           accessMode
		startOffset int
		endOffset   int
	}{
		{
			program:     []int{9, 9},
			m:           immediateMode,
			startOffset: 0,
			endOffset:   9,
		},
		{
			program:     []int{9, 2, 3},
			m:           positionMode,
			startOffset: 0,
			endOffset:   3,
		},
		{
			program:     []int{9, -1},
			m:           relativeMode,
			startOffset: 1,
			endOffset:   10,
		},
		{
			// offset index is 3 = -1 1+-1 =0
			program:     []int{9, 2, 0, -1},
			m:           relativeMode,
			startOffset: 1,
			endOffset:   0,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			c := instruction{
				Code:   9,
				P1Mode: tc.m,
			}

			p := &Process{
				program:            tc.program,
				relativeBaseOffset: tc.startOffset,
			}

			p.adjustRelativeBase(0, c)
			assert.Equal(t, tc.endOffset, p.relativeBaseOffset)
		})
	}
}
