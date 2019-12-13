package twelve

import (
	"fmt"
	"strings"
	"time"
)

type body struct {
	X, Y, Z    int //Verify this is an int
	vX, vY, vZ int // Starts at 0
}

func (b *body) ApplyGravity(a *body) {
	// x
	if b.X < a.X {
		b.vX++
	} else if b.X > a.X {
		b.vX--
	}
	// y
	if b.Y < a.Y {
		b.vY++
	} else if b.Y > a.Y {
		b.vY--
	}
	// z
	if b.Z < a.Z {
		b.vZ++
	} else if b.Z > a.Z {
		b.vZ--
	}
}

func (b *body) ApplyVelocity() {
	b.X += b.vX
	b.Y += b.vY
	b.Z += b.vZ
}

func (b *body) TotalEnergy() int {
	p := b.PotentialEnergy()
	k := b.KineticEnergy()

	return p * k
}

func (b *body) PotentialEnergy() int {
	return abs(b.X) + abs(b.Y) + abs(b.Z)
}
func (b *body) KineticEnergy() int {
	return abs(b.vX) + abs(b.vY) + abs(b.vZ)
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (b *body) String() string {
	return fmt.Sprintf("<%d,%d,%d>-<%d,%d,%d>", b.X, b.Y, b.Z, b.vX, b.vY, b.vZ)
}

func key(bodies []*body) string {
	sb := strings.Builder{}

	for _, b := range bodies {
		sb.WriteString(b.String())
	}

	return sb.String()
}

func FindRepeat(bodies []*body) int {
	startTime := time.Now()
	states := make(map[string]bool)
	states[key(bodies)] = true
	for i := 1; true; i++ {
		oneStep(bodies)
		k := key(bodies)

		if _, ok := states[k]; ok {
			return i
		}

		if i == 1000000 {
			fmt.Printf("Took %s to get to 1m\n", time.Now().Sub(startTime).String())
		}

		if i > 4686774924 {
			panic("We missed it!")
		}

		states[k] = true
	}

	return -1
}

func oneStep(bodies []*body) {
	// Apply gravity
	for _, b := range bodies {

		for _, a := range bodies {
			b.ApplyGravity(a)
		}

	}

	// Apply velocity
	for _, b := range bodies {
		b.ApplyVelocity()
	}

}

func SimulateNSteps(steps int, bodies []*body) {
	for i := 0; i < steps; i++ {
		oneStep(bodies)
	}
}

func TotalEnergy(bodies []*body) int {
	t := 0
	for _, b := range bodies {
		t += b.TotalEnergy()
	}

	return t
}
