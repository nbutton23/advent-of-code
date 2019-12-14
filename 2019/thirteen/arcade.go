package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/nbutton23/advent-of-code/2019/intcode"
)

const (
	empty = iota
	wall
	block
	horizontalPaddle
	ball
)

type Game struct {
	tiles [][]int

	inputChan  chan int //Channel that input to the robot is from
	outputChan chan int //Channel the robot write to

	process *intcode.Process // Intcode process
	program []int

	score int

	emptyCount            int
	wallCount             int
	blockCount            int
	horizontalPaddleCount int
	ballCount             int

	w, h int
}

func NewGame(program []int, x, y int) *Game {

	tiles := genEmptyBoard(x, y)

	inputChan := make(chan int)
	outputChan := make(chan int, 10)

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

	painter := &Game{
		tiles:      tiles,
		inputChan:  inputChan,
		outputChan: outputChan,

		process: process,
		program: program,

		w: x,
		h: y,
	}

	return painter
}

func controlerInput(output chan int) {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		if string(b) == "s" {
			output <- -1
		} else if string(b) == "d" {
			output <- 1
		} else if string(b) == "w" {
			output <- 0
		}

	}
}

func genEmptyBoard(maxX, maxY int) [][]int {

	hull := make([][]int, 0)
	for y := 0; y < maxY; y++ {
		row := make([]int, 0)

		for x := 0; x < maxX; x++ {
			row = append(row, empty)
		}

		hull = append(hull, row)
	}

	return hull
}

func writeScore(s int) {
	f, _ := os.Create("score.txt")

	f.WriteString(strconv.Itoa(s))
}

func (g *Game) Run() {
	g.process.IsHalted = false
	go g.process.ProcessProgram(0, g.program)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {

		paddelX := 0
		ballX := 0
		i := 0
		// always := false
		for !g.process.IsHalted {
			i++
			x := <-g.inputChan
			y := <-g.inputChan
			title := <-g.inputChan

			if x == -1 {
				writeScore(title)
				continue
			}

			g.tiles[y][x] = title

			switch title {
			case empty:
				g.emptyCount++
			case wall:
				g.wallCount++
			case block:
				g.blockCount++
			case horizontalPaddle:
				g.horizontalPaddleCount++
				paddelX = x
			case ball:
				g.ballCount++
				ballX = x
				if paddelX < ballX {
					g.outputChan <- 1
				} else if paddelX > ballX {
					g.outputChan <- -1
				} else {
					g.outputChan <- 0
				}
			}

			// time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		t := time.Tick(10 * time.Millisecond)
		for {
			select {
			case <-t:
				g.printBoard()
			}
		}
	}()

	// go controlerInput(g.outputChan)

	wg.Wait()

}

func (p *Game) printBoard() {
	scale := 20
	w, h := p.w, p.h
	rect := image.Rect(0, 0, w*scale, h*scale)
	img := image.NewRGBA(rect)
	green := color.RGBA{14, 184, 14, 0xff}
	brown := color.RGBA{178, 34, 34, 0xff}
	gray := color.RGBA{128, 128, 128, 0xff}
	blue := color.RGBA{30, 144, 255, 0xff}
	for r, row := range p.tiles {
		for c, p := range row {
			for y := r * scale; y < (r+1)*scale; y++ {
				for x := c * scale; x < (c+1)*scale; x++ {
					switch p {
					case empty:
						img.Set(x, y, color.Black)
					case wall:
						img.Set(x, y, gray)
					case block:
						img.Set(x, y, brown)
					case horizontalPaddle:
						img.Set(x, y, blue)
					case ball:
						img.Set(x, y, green)
					}
				}
			}
		}
	}
	f, _ := os.Create(fmt.Sprintf("img-%d-%d.png", w, h))
	png.Encode(f, img)
}
