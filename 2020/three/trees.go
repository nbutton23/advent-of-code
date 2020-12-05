package three

import "strings"

type slopePair struct {
	x, y int
}
type slop struct {
	// The pattern is an infint grid in the x
	pattern [][]string
	xLength int
	yLength int
}

func parsePatternIntoSlop(s string) slop {
	// Split by newline (rows)
	rows := strings.Split(s, "\n")
	// get length for yLength
	yLength := len(rows)
	// get length of row for xLength
	xLength := len(strings.TrimSpace(rows[0]))
	pattern := make([][]string, yLength)
	for i, r := range rows {

		r = strings.TrimSpace(r)
		row := make([]string, xLength)
		for j, c := range r {
			row[j] = string(c)
		}
		pattern[i] = row
	}

	return slop{
		pattern: pattern,
		xLength: xLength,
		yLength: yLength,
	}
}

func (s *slop) CountTreesInSlop(x, y int) int {
	treeCount := 0
	x1, y1 := 0, 0
	for {
		x1 += x
		y1 += y
		if s.isOffSlop(y1) {
			break
		}

		s := s.getSpot(x1, y1)
		if s == "#" {
			treeCount++
		}
	}
	return treeCount
}

func (s *slop) CountTreesInSlops(slops []slopePair) []int {
	counts := make([]int, len(slops))
	for i, p := range slops {
		c := s.CountTreesInSlop(p.x, p.y)
		counts[i] = c
	}

	return counts
}

func (s *slop) ProductOfTreesInSlops(slops []slopePair) int {
	c := s.CountTreesInSlops(slops)
	p := 1

	for _, i := range c {
		p = p * i
	}

	return p
}

func (s *slop) getSpot(x, y int) string {
	correctedX := s.correctX(x)

	if s.isOffSlop(y) {
		// TODO: handle this better
		return ""
	}

	return s.pattern[y][correctedX]
}

func (s *slop) correctX(x int) int {
	if x < s.xLength {
		return x
	}

	if x == s.xLength {
		return 0
	}

	r := x % s.xLength

	return r
}

func (s *slop) isOffSlop(y int) bool {
	return y >= s.yLength
}
