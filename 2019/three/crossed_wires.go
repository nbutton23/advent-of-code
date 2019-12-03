package three

import (
	// "sort"

	"strconv"
	"strings"
)

func FindDistanceToCross(wireOne, wireTwo string) int {
	wOne := PointsFromString(wireOne)
	wTwo := PointsFromString(wireTwo)
	crossPoints := FindCrossPoints(wOne, wTwo)
	shortestDist := -1
	for _, d := range crossPoints {
		if d.ManhattanDistance() < shortestDist || shortestDist == -1 {
			shortestDist = d.ManhattanDistance()
		}
	}

	return shortestDist
}

func FindCrossPoints(wOne, wTwo points) points {
	pointsThatCross := make([]point, 0)

	for _, p1 := range wOne {
		for _, p2 := range wTwo {

			if p1.Equal(p2) {
				pointsThatCross = append(pointsThatCross, p1)
			}
		}
	}

	for i, j := 0, 0; i < len(wOne) && j < len(wTwo); {
		p1 := wOne[i]
		p2 := wTwo[j]
		if p1.Equal(p2) {
			pointsThatCross = append(pointsThatCross, p1)
			i++
			j++
		} else if p1.LessThan(p2) {
			i++
		} else {
			j++
		}
	}

	return pointsThatCross
}

func FindBestSteps(wireOne string, wireTwo string) int {
	wOne := PointsFromString(wireOne)
	wTwo := PointsFromString(wireTwo)
	crossPoints := FindCrossPoints(wOne, wTwo)

	shortest := -1

	for _, p := range crossPoints {
		stepsToOne := FindStepsToPoint(wOne, p)
		stepsToTwo := FindStepsToPoint(wTwo, p)

		// Plus 2 because i am skipping the first step from 0,0
		total := stepsToOne + stepsToTwo + 2

		if total < shortest || shortest == -1 {
			shortest = total
		}
	}

	return shortest
}

func FindStepsToPoint(points []point, pointToFind point) int {
	for i, p := range points {
		if p.Equal(pointToFind) {
			return i
		}
	}

	return -1
}

func PointsFromString(wire string) points {

	dirs := strings.Split(wire, ",")
	points := make([]point, 0)
	x := 0
	y := 0

	for _, instruction := range dirs {
		direction := instruction[0:1]
		dstring := instruction[1:]
		distance, err := strconv.Atoi(dstring)
		if err != nil {
			// TODO
			panic(err)
		}
		switch direction {
		case "R":
			for i := 0; i < distance; i++ {
				x++
				points = append(points, point{y: y, x: x, char: '-'})
			}
			points[len(points)-1].char = '+'
		case "L":
			for i := 0; i < distance; i++ {
				x--
				points = append(points, point{y: y, x: x, char: '-'})
			}
			points[len(points)-1].char = '+'
		case "U":
			for i := 0; i < distance; i++ {
				y++
				points = append(points, point{y: y, x: x, char: '|'})
			}
			points[len(points)-1].char = '+'
		case "D":
			for i := 0; i < distance; i++ {
				y--
				points = append(points, point{y: y, x: x, char: '|'})
			}
			points[len(points)-1].char = '+'
		default:
			// TODO
			panic(direction)
		}
	}

	return points
}

type point struct {
	x, y int
	char rune
}

func (p point) Equal(in point) bool {
	return in.x == p.x && in.y == p.y
}

func (p point) LessThan(in point) bool {
	return p.x < in.x
}

func (p point) ManhattanDistance() int {
	absX := p.x
	if absX < 0 {
		absX *= -1
	}

	absY := p.y
	if absY < 0 {
		absY *= -1
	}

	return absX + absY
}

type points []point

func (p points) Len() int {
	return len(p)
}
func (p points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p points) Less(i, j int) bool {
	return p[i].ManhattanDistance() < p[j].ManhattanDistance()
}

func (p points) Print() string {
	xMax := -1
	yMax := -1
	xMin := -1
	yMin := -1

	for _, r := range p {
		if r.x > xMax {
			xMax = r.x
		}
		if r.y > yMax {
			yMax = r.y
		}
		if r.x < xMin {
			xMin = r.x
		}
		if r.y < yMin {
			yMin = r.y
		}
	}
	xMax += 2
	yMax += 2

	xMax = (xMin * -1) + xMax
	yMax = (yMin * -1) + yMax

	graph := make([][]rune, yMax)

	for i := 0; i < yMax; i++ {
		graph[i] = make([]rune, xMax)
		for j := 0; j < xMax; j++ {
			// set the my zero value to *
			graph[i][j] = '*'
		}
	}
	for _, r := range p {
		graph[r.y-yMin][r.x-xMin] = r.char
	}

	graph[0-yMin][0-xMin] = 'O'

	var sb strings.Builder
	for i := yMax - 1; i >= 0; i-- {
		// for _, row := range graph {
		sb.WriteString(string(graph[i]))
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (p points) PrintWithOverlay(o points) string {
	xMax := -1
	yMax := -1
	xMin := -1
	yMin := -1

	for _, r := range p {
		if r.x > xMax {
			xMax = r.x
		}
		if r.y > yMax {
			yMax = r.y
		}
		if r.x < xMin {
			xMin = r.x
		}
		if r.y < yMin {
			yMin = r.y
		}
	}

	for _, r := range o {
		if r.x > xMax {
			xMax = r.x
		}
		if r.y > yMax {
			yMax = r.y
		}
		if r.x < xMin {
			xMin = r.x
		}
		if r.y < yMin {
			yMin = r.y
		}
	}
	xMax += 2
	yMax += 2

	xMax = (xMin * -1) + xMax
	yMax = (yMin * -1) + yMax

	graph := make([][]rune, yMax)

	for i := 0; i < yMax; i++ {
		graph[i] = make([]rune, xMax)
		for j := 0; j < xMax; j++ {
			// set the my zero value to *
			graph[i][j] = '*'
		}
	}
	for _, r := range p {

		if graph[r.y-yMin][r.x-xMin] != '*' {
			graph[r.y-yMin][r.x-xMin] = 'x'
		} else {
			graph[r.y-yMin][r.x-xMin] = r.char
		}
	}
	for _, r := range o {

		if graph[r.y-yMin][r.x-xMin] != '*' {
			graph[r.y-yMin][r.x-xMin] = 'x'
		} else {
			graph[r.y-yMin][r.x-xMin] = r.char
		}
	}

	graph[0-yMin][0-xMin] = 'O'

	var sb strings.Builder
	for i := yMax - 1; i >= 0; i-- {
		// for _, row := range graph {
		sb.WriteString(string(graph[i]))
		sb.WriteRune('\n')
	}

	return sb.String()
}
