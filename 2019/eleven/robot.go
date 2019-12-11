package eleven

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/nbutton23/advent-of-code/2019/intcode"
)

type direction string

const (
	north = direction("^") // turn Right is +x turn Left is -x
	south = direction("v") // turn Right is -x turn Left is +x
	west  = direction("<") // turn Right is -y turn Left is +y
	east  = direction(">") // turn right is +y turn left is -y
)

type panel struct {
	color       int
	beenPainted bool
}
type Painter struct {
	hull             [][]panel
	x, y             int // position on the hull
	currentDirection direction

	inputChan  chan int //Channel that input to the robot is from
	outputChan chan int //Channel the robot write to

	process               *intcode.Process // Intcode process
	numberOfPanelsPainted int
	program               []int
}

func NewPainter(program []int, x, y int) *Painter {

	hull := genHull(100, 100)

	inputChan := make(chan int)
	outputChan := make(chan int)

	pOstream := func(a ...interface{}) (n int, err error) {
		o, ok := a[0].(int)
		if !ok {
			panic("expected int")
		}
		inputChan <- o
		return 0, nil
	}

	pInStream := func() int {
		return <-outputChan
	}

	process := intcode.NewProccess(pInStream, pOstream)

	painter := &Painter{
		hull:             hull,
		inputChan:        inputChan,
		outputChan:       outputChan,
		currentDirection: north,
		x:                x,
		y:                y,
		process:          process,
		program:          program,
	}

	return painter
}

func genHull(maxX, maxY int) [][]panel {

	hull := make([][]panel, 0)
	for y := 0; y < maxY; y++ {
		row := make([]panel, 0)

		for x := 0; x < maxX; x++ {
			color := 0

			row = append(row, panel{
				color:       color,
				beenPainted: false,
			})
		}

		hull = append(hull, row)
	}

	return hull
}
func (p *Painter) Close() {
	close(p.inputChan)
	close(p.inputChan)

	p.process = nil
	p = nil
}
func (p *Painter) Run() int {
	//Start process
	p.process.IsHalted = false
	go p.process.ProcessProgram(0, p.program)

	for !p.process.IsHalted {
		panel := p.hull[p.y][p.x]
		//Read panel color and push to outputchan
		p.outputChan <- panel.color
		// Read two values off of input
		newColor := <-p.inputChan
		newDir := <-p.inputChan
		// if 0 paint black if 1 paint white
		panel.color = newColor
		if panel.beenPainted == false {
			p.numberOfPanelsPainted++
			panel.beenPainted = true
		}
		p.hull[p.y][p.x] = panel
		p.turnBot(newDir)
	}
	p.printHull()
	return p.numberOfPanelsPainted
}

func (p *Painter) turnBot(dir int) {

	turnRight := 1
	switch p.currentDirection {
	case north:
		// turn Right(>) (1) is +x turn Left(<) (0) is -x
		if dir == turnRight {
			p.currentDirection = east
			p.x++
		} else {
			p.currentDirection = west
			p.x--
		}
	case south:
		// turn Right (<) is -x turn Left(>) is +x
		if dir == turnRight {
			p.currentDirection = west
			p.x--
		} else {
			p.currentDirection = east
			p.x++
		}
	case east:
		// turn right (v) is ++y turn left(^) is --y
		if dir == turnRight {
			p.currentDirection = south
			p.y++
		} else {
			p.currentDirection = north
			p.y--
		}
	case west:
		// turn right(^) is --y turn left(v) is ++y
		if dir == turnRight {
			p.currentDirection = north
			p.y--
		} else {
			p.currentDirection = south
			p.y++
		}
	}

	if p.x < 0 || p.x >= len(p.hull) || p.y < 0 || p.y >= len(p.hull) {
		p.printHull()
		panic("Will out of bounds! Printing board")
	}
}

func (p *Painter) printHull() {
	w, h := 100, 100
	rect := image.Rect(0, 0, w*20, h*20)
	img := image.NewRGBA(rect)
	green := color.RGBA{14, 184, 14, 0xff}
	for r, row := range p.hull {
		for c, p := range row {
			for y := r * 20; y < (r+1)*20; y++ {
				for x := c * 20; x < (c+1)*20; x++ {
					switch p.color {
					case 0:
						img.Set(x, y, color.Black)
					case 1:
						img.Set(x, y, green)
					}
				}
			}
		}
	}
	f, _ := os.Create(fmt.Sprintf("img-%d-%d.png", w, h))
	png.Encode(f, img)
}
