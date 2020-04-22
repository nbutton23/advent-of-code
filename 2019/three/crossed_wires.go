package three

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

func FindDistanceToCross(wireOne, wireTwo string) int {
	wOne := PointsFromString(wireOne)
	wTwo := PointsFromString(wireTwo)
	crossPoints := FindCrossPointsV2(wOne, wTwo)
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

	return pointsThatCross
}

func FindCrossPointsV2(wOne, wTwo points) points {

	wOneMap := make(map[string]point)

	for _, p := range wOne {
		wOneMap[p.Key()] = p
	}

	wTwoMap := make(map[string]point)

	for _, p := range wTwo {
		wTwoMap[p.Key()] = p
	}

	pointsThatCross := make([]point, 0)
	for k, _ := range wOneMap {
		if mp, ok := wTwoMap[k]; ok {
			mp.char = 'x'
			pointsThatCross = append(pointsThatCross, mp)
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

func (p point) Key() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
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

func (p points) minMax() (xMin, yMin, xMax, yMax int) {

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

	return
}

func (p points) fillInImage(img *image.RGBA, color color.Color, xMin, yMin int, blockSize int) {
	for _, r := range p {
		for y := (r.y - yMin) * blockSize; y < (r.y-yMin+1)*blockSize; y++ {
			for x := (r.x - xMin) * blockSize; x < (r.x-xMin+1)*blockSize; x++ {
				img.Set(x, y, color)
			}
		}
	}
}
func (p points) PNG() string {
	xMin, yMin, xMax, yMax := p.minMax()

	xMax += 2
	yMax += 2

	xMax = (xMin * -1) + xMax
	yMax = (yMin * -1) + yMax

	green := color.RGBA{14, 184, 14, 0xff}
	squarSiz := 20
	rect := image.Rect(0, 0, xMax*squarSiz, yMax*squarSiz)
	img := image.NewRGBA(rect)

	p.fillInImage(img, green, xMin, yMin, squarSiz)

	f, _ := os.Create(fmt.Sprintf("image-%d-%d.png", xMax, yMax))

	png.Encode(f, img)
	return f.Name()
}

func (p points) PNGWithOverlay(o points) string {
	xMin, yMin, xMax, yMax := p.minMax()
	oxMin, oyMin, oxMax, oyMax := p.minMax()

	if oxMin < xMin {
		xMin = oxMin
	}

	if oyMin < yMin {
		yMin = oyMin
	}

	if oxMax > xMax {
		xMax = oxMax
	}

	if oyMax > yMax {
		yMax = oyMax
	}

	xMax += 2
	yMax += 2

	xMax = (xMin * -1) + xMax
	yMax = (yMin * -1) + yMax

	squarSiz := 2
	rect := image.Rect(0, 0, xMax*squarSiz, yMax*squarSiz)
	img := image.NewRGBA(rect)

	green := color.RGBA{14, 184, 14, 0xff}
	blue := color.RGBA{0, 0, 0xff, 0xff}

	p.fillInImage(img, green, xMin, yMin, squarSiz)
	o.fillInImage(img, blue, xMin, yMin, squarSiz)

	f, _ := os.Create(fmt.Sprintf("overlay-%d-%d.png", xMax, yMax))

	png.Encode(f, img)
	return f.Name()
}
