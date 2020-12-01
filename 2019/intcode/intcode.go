package intcode

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var defaultInStream = func() int {
	bytes, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	text := string(bytes)
	inInt, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return inInt
}

type Process struct {
	inStream       func() int
	oStream        func(a ...interface{}) (n int, err error)
	instructionSet map[int]func(i int, c instruction) int

	program            []int
	outputs            []int
	relativeBaseOffset int

	IsHalted bool
}

func NewProccess(inStream func() int, oStream func(a ...interface{}) (n int, err error)) *Process {
	if inStream == nil {
		inStream = defaultInStream
	}

	if oStream == nil {
		oStream = fmt.Println
	}

	p := &Process{
		inStream: inStream,
		oStream:  oStream,

		outputs:            make([]int, 0),
		relativeBaseOffset: 0,
	}

	p.instructionSet = map[int]func(i int, c instruction) int{
		1:  p.add,
		2:  p.multiply,
		3:  p.input,
		4:  p.output,
		5:  p.jumpIfTrue,
		6:  p.jumpIfFalse,
		7:  p.lessThan,
		8:  p.equal,
		9:  p.adjustRelativeBase,
		99: p.halt,
	}

	return p
}

// ProcessProgram takes a intcode program and runs it to completion or error
func (p *Process) ProcessProgram(i int, program []int) error {
	p.IsHalted = false

	p.program = program
	for i < len(p.program) && i >= 0 {
		code := p.getValue(i)
		c := BreakUpOpCode(code)
		if cmd, ok := p.instructionSet[c.Code]; ok {
			i = cmd(i, c)
		} else {
			return fmt.Errorf("unknown instruction: %v", c)
		}
	}
	p.IsHalted = true

	return nil
}

type accessMode int

const (
	positionMode  = accessMode(0)
	immediateMode = accessMode(1)
	relativeMode  = accessMode(2)
)

func getAccessMode(m int) accessMode {
	switch m {
	case 0:
		return positionMode
	case 1:
		return immediateMode
	case 2:
		return relativeMode
	default:
		panic(fmt.Sprintf("Unknown Access Mode %d", m))
	}
}

type instruction struct {
	Code   int
	P1Mode accessMode
	P2Mode accessMode
	P3Mode accessMode
}

func BreakUpOpCode(code int) instruction {

	c := ((digit(code, 2) * 10) + digit(code, 1))
	p1 := getAccessMode(digit(code, 3))
	p2 := getAccessMode(digit(code, 4))
	p3 := getAccessMode(digit(code, 5))

	op := instruction{
		Code:   c,
		P1Mode: p1,
		P2Mode: p2,
		P3Mode: p3,
	}
	return op
}

// https://stackoverflow.com/questions/46753736/extract-digits-at-a-certain-position-in-a-number
func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
