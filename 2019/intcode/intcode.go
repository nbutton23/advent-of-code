package intcode

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
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
	cmdsRan        bytes.Buffer
	decomp         bool
	instructionSet map[int]func(i int, c instruction, program []int) int

	program            []int
	outputs            []int
	relativeBaseOffset int
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

		cmdsRan: bytes.Buffer{},
		decomp:  true,

		outputs:            make([]int, 0),
		relativeBaseOffset: 0,
	}

	p.instructionSet = map[int]func(i int, c instruction, program []int) int{
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
	p.program = program
	for i < len(program) && i >= 0 {
		c := BreakUpOpCode(program[i])
		if cmd, ok := p.instructionSet[c.Code]; ok {
			i = cmd(i, c, program)
		} else {
			return fmt.Errorf("unknown instruction: %v", c)
		}
	}

	tmpfile, _ := ioutil.TempFile("temp", "program-output")
	tmpfile.Write(p.cmdsRan.Bytes())
	tmpfile.Close()
	return nil
}

func (p *Process) halt(i int, c instruction, program []int) int {
	if p.decomp {
		p.cmdsRan.WriteString("HALT\n")
	}
	return -1
}

func (p *Process) add(i int, c instruction, program []int) int {
	// Add [input1, input2, output]
	p1 := p.valueForPram(i+1, program, c.P1Mode)
	p2 := p.valueForPram(i+2, program, c.P2Mode)
	p3 := p.indexForPram(i+3, program, c.P3Mode)

	program[p3] = p1 + p2

	if p.decomp {
		pr1 := p.retPramString(i+1, c.P1Mode)
		pr2 := p.retPramString(i+2, c.P2Mode)
		pr3 := p.retPramString(i+3, immediateMode)

		p.cmdsRan.WriteString(fmt.Sprintf("ADD %s %s %s\n", pr1, pr2, pr3))
	}
	return i + 4
}

func (p *Process) multiply(i int, c instruction, program []int) int {
	p1 := p.valueForPram(i+1, program, c.P1Mode)
	p2 := p.valueForPram(i+2, program, c.P2Mode)
	p3 := p.indexForPram(i+3, program, c.P3Mode)

	program[p3] = p1 * p2

	if p.decomp {
		pr1 := p.retPramString(i+1, c.P1Mode)
		pr2 := p.retPramString(i+2, c.P2Mode)
		pr3 := p.retPramString(i+3, immediateMode)

		p.cmdsRan.WriteString(fmt.Sprintf("MUL %s %s %s\n", pr1, pr2, pr3))
	}

	return i + 4
}

func (p *Process) input(i int, c instruction, program []int) int {
	// input store at i+1
	p1 := p.indexForPram(i+1, program, c.P1Mode)
	program[p1] = p.inStream()

	if p.decomp {
		pr1 := p.retPramString(i+1, immediateMode)

		p.cmdsRan.WriteString(fmt.Sprintf("IN %s\n", pr1))
	}

	return i + 2
}
func (p *Process) output(i int, c instruction, program []int) int {
	p1 := p.valueForPram(i+1, program, c.P1Mode)
	out := p1
	p.outputs = append(p.outputs, out)
	p.oStream(out)
	if p.decomp {
		pr1 := p.retPramString(i+1, c.P1Mode)

		p.cmdsRan.WriteString(fmt.Sprintf("OUT %s - %d\n", pr1, out))
	}
	return i + 2
}

func (p *Process) jumpIfTrue(i int, c instruction, program []int) int {
	// jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value
	// from the second parameter. Otherwise, it does nothing.
	p1 := p.valueForPram(i+1, program, c.P1Mode)

	if p1 != 0 {
		p2 := p.valueForPram(i+2, program, c.P2Mode)
		return p2

	}
	if p.decomp {
		pr1 := p.retPramString(i+1, c.P1Mode)
		pr2 := p.retPramString(i+2, c.P2Mode)

		p.cmdsRan.WriteString(fmt.Sprintf("JEQ %s %s\n", pr1, pr2))
	}
	return i + 3

}
func (p *Process) jumpIfFalse(i int, c instruction, program []int) int {
	// jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value
	// from the second parameter. Otherwise, it does nothing.
	p1 := p.valueForPram(i+1, program, c.P1Mode)

	if p1 == 0 {
		p2 := p.valueForPram(i+2, program, c.P2Mode)
		return p2

	}
	if p.decomp {
		pr1 := p.retPramString(i+1, c.P1Mode)
		pr2 := p.retPramString(i+2, c.P2Mode)

		p.cmdsRan.WriteString(fmt.Sprintf("JNE %s %s\n", pr1, pr2))
	}
	return i + 3

}

func (p *Process) lessThan(i int, c instruction, program []int) int {
	// less than: if the first parameter is less than the second parameter,
	// it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	p1 := p.valueForPram(i+1, program, c.P1Mode)
	p2 := p.valueForPram(i+2, program, c.P2Mode)
	p3 := p.indexForPram(i+3, program, c.P3Mode) //Store at the position returned from here

	if p1 < p2 {
		program[p3] = 1
	} else {
		program[p3] = 0
	}

	if p.decomp {
		pr1 := p.retPramString(i+1, c.P1Mode)
		pr2 := p.retPramString(i+2, c.P2Mode)
		pr3 := p.retPramString(i+3, immediateMode)

		p.cmdsRan.WriteString(fmt.Sprintf("LES %s %s %s\n", pr1, pr2, pr3))
	}

	return i + 4
}

func (p *Process) equal(i int, c instruction, program []int) int {
	// equals: if the first parameter is equal to the second parameter,
	// it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	p1 := p.valueForPram(i+1, program, c.P1Mode)
	p2 := p.valueForPram(i+2, program, c.P2Mode)
	p3 := p.indexForPram(i+3, program, c.P3Mode)

	if p1 == p2 {
		program[p3] = 1
	} else {
		program[p3] = 0
	}

	if p.decomp {
		pr1 := p.retPramString(i+1, c.P1Mode)
		pr2 := p.retPramString(i+2, c.P2Mode)
		pr3 := p.retPramString(i+3, positionMode)

		p.cmdsRan.WriteString(fmt.Sprintf("EQL %s %s %s\n", pr1, pr2, pr3))
	}
	return i + 4
}

func (p *Process) adjustRelativeBase(i int, c instruction, program []int) int {
	// adjusts the relative base by the value of its only parameter.
	// The relative base increases (or decreases, if the value is negative)
	// by the value of the parameter.
	p1 := p.valueForPram(i+1, program, c.P1Mode)

	p.relativeBaseOffset += p1
	return i + 1
}

// TODO: Create Getters and Setters that will allow us to expand pass the current memory
func (p *Process) setValue(index int, value int, program []int) {
	if index < 0 {
		panic("Negitive index!")
	}

	if index >= len(program) {

	}
}

// Returns the Value a value.
func (p *Process) valueForPram(i int, program []int, mode accessMode) int {
	switch mode {
	case immediateMode:
		return program[i]
	case positionMode:
		return program[program[i]]
	case relativeMode:
		pos := p.relativeBaseOffset - program[i]
		return program[pos]
	default:
		panic("Received unknown mode")
	}
}

// Returns the index
func (p *Process) indexForPram(i int, program []int, mode accessMode) int {
	switch mode {
	case immediateMode:
		return i
	case positionMode:
		return program[i]
	case relativeMode:
		pos := p.relativeBaseOffset - program[i]
		return pos
	default:
		panic("Received unknown mode")
	}
}

func (p *Process) retPramString(i int, mode accessMode) string {
	switch mode {
	case immediateMode:
		return fmt.Sprintf("*%d", i)
	case positionMode:
		return fmt.Sprintf("%d", i)
	default:
		panic("Received unknown mode")
	}
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
		panic("unknown accessMode type")
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
