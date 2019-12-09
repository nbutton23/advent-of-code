package eight

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetLayers(t *testing.T) {
	loadTestData()
	testCases := []struct {
		input    string
		w, h     int
		expected [][]string
	}{
		{
			input: "123456789012",
			w:     3,
			h:     2,
			expected: [][]string{
				[]string{"123", "456"},
				[]string{"789", "012"},
			},
		},
		// {
		// 	input: testInput,
		// 	w:     25,
		// 	h:     6,
		// 	// Invalid
		// 	expected: [][]string{
		// 		[]string{"123", "456"},
		// 		[]string{"789", "012"},
		// 	},
		// },
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			layers := GetLayers(tc.input, tc.w, tc.h)
			require.Len(t, layers, len(tc.expected))
			for i := 0; i < len(tc.expected); i++ {
				assert.ElementsMatch(t, layers[i], tc.expected[i])
			}
		})
	}
}

func Test_RenderImage(t *testing.T) {
	loadTestData()

	testCases := []struct {
		input    string
		w, h     int
		expected string
	}{
		{
			input:    "0222112222120000",
			w:        2,
			h:        2,
			expected: "0110",
		},
		{
			input:    testInput,
			w:        25,
			h:        6,
			expected: "100000110010001100101110010000100101000110010100101000010000010101111011100100001011000100100101001010000100100010010010100101111001110001001001011100",
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			out := RenderImage(tc.input, tc.w, tc.h)
			
			assert.Equal(t, tc.expected, out)
		})
	}
}
func Test_ImageCheckSum(t *testing.T) {
	loadTestData()

	testCases := []struct {
		input    string
		w, h     int
		expected int
	}{
		{
			input:    testInput,
			w:        25,
			h:        6,
			expected: 1088,
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			check := ImageCheckSum(tc.input, tc.w, tc.h)
			assert.Equal(t, tc.expected, check)
		})
	}
}
func Test_CountNumberLayer(t *testing.T) {
	testCases := []struct {
		input []string
		cut   string
		ex    int
	}{
		{
			[]string{"121", "456"},
			"1",
			2,
		},
		{
			[]string{"020", "456"},
			"1",
			0,
		},
		{
			[]string{"121", "416"},
			"1",
			3,
		},
		{
			[]string{"111", "111"},
			"1",
			6,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			c := CountNumberLayer(tc.input, tc.cut)
			assert.Equal(t, tc.ex, c)
		})
	}
}

func Test_LayerWithMax(t *testing.T) {
	testCases := []struct {
		input [][]string
		cut   string
		ex    int
	}{
		{
			[][]string{
				[]string{"121", "456"},
				[]string{"020", "456"},
			},
			"1",
			0,
		},
		{
			[][]string{
				[]string{"020", "456"},
				[]string{"121", "456"},
			},
			"1",
			1,
		},
		{
			[][]string{
				[]string{"121", "156"},
				[]string{"120", "156"},
			},
			"1",
			0,
		},
		{
			[][]string{
				[]string{"120", "156"},
				[]string{"121", "156"},
			},
			"1",
			1,
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			c := LayerWithMax(tc.input, tc.cut)
			assert.Equal(t, tc.ex, c)
		})
	}
}
