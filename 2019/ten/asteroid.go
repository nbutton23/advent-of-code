package ten

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Astroid struct {
	X, Y int

	canSee map[float64][]*Astroid
}

func NewAstroid(x, y int) *Astroid {
	return &Astroid{
		X:      x,
		Y:      y,
		canSee: make(map[float64][]*Astroid),
	}
}
func (a *Astroid) ID() string {
	return fmt.Sprintf("%d,%d", a.X, a.Y)
}

func (a *Astroid) LineOfSiteCount() int {
	return len(a.canSee)
}

func (a *Astroid) CheckIfLineOfSite(astroids []*Astroid) {
	for _, b := range astroids {
		if a.X == b.X && a.Y == b.Y {
			continue
		}
		angle := a.BearingToAstroid(b)

		if _, ok := a.canSee[angle]; !ok {
			a.canSee[angle] = make([]*Astroid, 0)
		}

		a.canSee[angle] = append(a.canSee[angle], b)
	}
}

/*
Returns the angle counter clockwise from "north"
*/
func (a *Astroid) BearingToAstroid(b *Astroid) float64 {

	x := (a.X - b.X)
	y := (a.Y - b.Y)

	gcd := GCD(x, y)
	if gcd < 0 {
		// make it positive
		gcd = -gcd
	} else if gcd == 0 {
		gcd = 1
	}

	x = x / gcd
	y = y / gcd
	angle := math.Atan2(float64(y), float64(x))/(2*math.Pi)*360.0 + 90
	if angle < 0 {
		angle += 360
	}

	return angle
}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func (a *Astroid) ReturnNthAstroidToDestroy(n int) *Astroid {
	a.Sort()

	keys := make([]float64, 0, len(a.canSee))
	for k := range a.canSee {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		// Start at 180 and go up until we wrap backaround at 360
		vI := keys[i]
		vJ := keys[j]
		if vI < 180 && vJ < 180 {
			return vI < vJ
		} else if vI >= 180 && vJ >= 180 {
			return vI < vJ
		} else if vI <= 180 {
			return false
		}
		return true
	})

	i := 0
	// this is destructive to the backing maps!
	for i < n {
		for _, k := range keys {
			r := a.canSee[k]
			if len(r) > 0 {
				i++
				if i == n {
					return r[0]
				}
				// clear off one
				// fmt.Printf("The %d asteroid to be vaporized is at %s\n", i, r[0].ID())
				a.canSee[k] = r[1:]
			}
		}
	}

	return nil
}
func (a *Astroid) DistanceToAstroid(b *Astroid) float64 {
	// 	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))

	x := float64(a.X - b.X)
	y := float64(a.Y - b.Y)

	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func (a *Astroid) Sort() {
	for _, an := range a.canSee {

		sort.Slice(an, func(i, j int) bool {
			return an[i].DistanceToAstroid(a) < an[j].DistanceToAstroid(a)
		})
	}
}

func FindAstroidWithMaxSight(astroids []*Astroid) *Astroid {
	max := -1
	var astroid *Astroid
	for _, a := range astroids {
		a.CheckIfLineOfSite(astroids)
		c := a.LineOfSiteCount()
		if c > max {
			max = c
			astroid = a
		}
	}

	return astroid
}

func GetAstroids(astroidMap string) []*Astroid {
	astroids := make([]*Astroid, 0)
	rows := strings.Split(astroidMap, "\n")
	if len(rows) == 0 {
		return nil
	}
	// making the assumption that the data is well formed
	cols := len(rows[0])
	for y, r := range rows {
		for x := 0; x < cols; x++ {
			if r[x] == '#' {
				astroids = append(astroids, NewAstroid(x, y))
			}
		}
	}

	return astroids
}
