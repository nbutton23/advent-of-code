package two

import "testing"

import "github.com/stretchr/testify/assert"

func Test_ProccessInstruction(t *testing.T) {
	p := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

	err := ProcessProgram(0, p)

	assert.NoError(t, err)

	assert.Equal(t, p[0], 3500)
}

func Test_ProcessGravityAssist(t *testing.T) {
	o := ProcessGravityAssist(12, 2)

	assert.Equal(t, 3765464, o)
}

func Test_FindNounAndVerbForOutput(t *testing.T) {
	testCases := []struct {
		output       int
		min          int
		max          int
		expectedNoun int
		expectedVerb int
	}{
		{3765464, 0, 99, 12, 2},
		{19690720, 0, 99, 76, 10},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			n, v := FindNounAndVerbForOutput(tc.output, tc.min, tc.max)
			assert.Equal(t, tc.expectedNoun, n)
			assert.Equal(t, tc.expectedVerb, v)
		})
	}
}
