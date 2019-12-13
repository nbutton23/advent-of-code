package twelve

import "testing"

import "github.com/stretchr/testify/assert"

func Test_PotentialEnergy(t *testing.T) {
	testCases := []struct {
		x, y, z int
		exp     int
	}{
		{
			2, 1, -3,
			6,
		},
		{
			1, -8, 0,
			9,
		},
		{
			3, -6, 1,
			10,
		},
		{
			2, 0, 4,
			6,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			b := body{
				X: tc.x,
				Y: tc.y,
				Z: tc.z,
			}

			p := b.PotentialEnergy()
			assert.Equal(t, tc.exp, p)
		})
	}
}

func Test_KineticEnergy(t *testing.T) {
	testCases := []struct {
		x, y, z int
		exp     int
	}{
		{
			-3, -2, 1,
			6,
		},
		{
			-1, 1, 3,
			5,
		},
		{
			3, 2, -3,
			8,
		},
		{
			1, -1, -1,
			3,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			b := body{
				vX: tc.x,
				vY: tc.y,
				vZ: tc.z,
			}

			p := b.KineticEnergy()
			assert.Equal(t, tc.exp, p)
		})
	}
}

func Test_TotalEnergy(t *testing.T) {
	testCases := []struct {
		x, y, z    int
		vX, vY, vZ int

		exp int
	}{
		{
			2, 1, -3,
			-3, -2, 1,
			36,
		},
		{
			1, -8, 0,
			-1, 1, 3,
			45,
		},
		{
			3, -6, 1,
			3, 2, -3,
			80,
		},
		{
			2, 0, 4,
			1, -1, -1,
			18,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			b := body{
				X: tc.x,
				Y: tc.y,
				Z: tc.z,

				vX: tc.vX,
				vY: tc.vY,
				vZ: tc.vZ,
			}

			p := b.TotalEnergy()
			assert.Equal(t, tc.exp, p)
		})
	}
}
func Test_Ex1(t *testing.T) {
	body1 := &body{
		X: -1, Y: 0, Z: 2,
	}
	body2 := &body{
		X: 2, Y: -10, Z: -7,
	}
	body3 := &body{
		X: 4, Y: -8, Z: 8,
	}
	body4 := &body{
		X: 3, Y: 5, Z: -1,
	}

	bodies := []*body{body1, body2, body3, body4}

	oneStep(bodies)
	testBodyVals(t, 2, -1, 1, 3, -1, -1, bodies[0])
}
func testBodyVals(t *testing.T, x, y, z, vX, vY, vZ int, b *body) {
	assert.Equal(t, x, b.X, "X")
	assert.Equal(t, y, b.Y, "Y")
	assert.Equal(t, z, b.Z, "Z")

	assert.Equal(t, vX, b.vX, "vX")
	assert.Equal(t, vY, b.vY, "vY")
	assert.Equal(t, vZ, b.vZ, "vZ")
}
func Test_Ex2(t *testing.T) {
	body1 := &body{
		X: -8, Y: -10, Z: 0,
	}
	body2 := &body{
		X: 5, Y: 5, Z: 10,
	}
	body3 := &body{
		X: 2, Y: -7, Z: 3,
	}
	body4 := &body{
		X: 9, Y: -8, Z: -3,
	}

	bodies := []*body{body1, body2, body3, body4}

	SimulateNSteps(100, bodies)
	total := TotalEnergy(bodies)
	assert.Equal(t, 1940, total)
}
func Test_Part1(t *testing.T) {
	body1 := &body{
		X: 1, Y: 4, Z: 4,
	}
	body2 := &body{
		X: -4, Y: -1, Z: 19,
	}
	body3 := &body{
		X: -15, Y: -14, Z: 12,
	}
	body4 := &body{
		X: -17, Y: 1, Z: 10,
	}

	bodies := []*body{body1, body2, body3, body4}

	SimulateNSteps(1000, bodies)
	total := TotalEnergy(bodies)
	assert.Equal(t, 10635, total)
}

func Test_Ex1_Repeat(t *testing.T) {
	body1 := &body{
		X: -1, Y: 0, Z: 2,
	}
	body2 := &body{
		X: 2, Y: -10, Z: -7,
	}
	body3 := &body{
		X: 4, Y: -8, Z: 8,
	}
	body4 := &body{
		X: 3, Y: 5, Z: -1,
	}

	bodies := []*body{body1, body2, body3, body4}

	c := FindRepeat(bodies)
	assert.Equal(t, 2772, c)
}

func Test_Ex2_Repeat(t *testing.T) {
	body1 := &body{
		X: -8, Y: -10, Z: 0,
	}
	body2 := &body{
		X: 5, Y: 5, Z: 10,
	}
	body3 := &body{
		X: 2, Y: -7, Z: 3,
	}
	body4 := &body{
		X: 9, Y: -8, Z: -3,
	}

	bodies := []*body{body1, body2, body3, body4}

	c := FindRepeat(bodies)
	assert.Equal(t, 4686774924, c)
}
