package intcode

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var oStream = fmt.Println
var instructionSet = map[int]func(i int, c intermediatInstruct, program []int) int{
	1:  add,
	2:  multiply,
	3:  input,
	4:  output,
	5:  jumpIfTrue,
	6:  jumpIfFalse,
	7:  lessThan,
	8:  equal,
	99: halt,
}

// ProcessProgram takes a intcode program and runs it to completion or error
func ProcessProgram(i int, program []int) error {

	for i < len(program) && i >= 0 {
		c := BreakUpOpCode(program[i])
		if cmd, ok := instructionSet[c.Code]; ok {
			i = cmd(i, c, program)
		} else {
			return fmt.Errorf("unknown instruction: %v", c)
		}
	}
	return nil
}

func halt(i int, c intermediatInstruct, program []int) int {
	return -1
}

func add(i int, c intermediatInstruct, program []int) int {
	// Add [input1, input2, output]
	p1 := retPram(i+1, program, c.P1Immediate)
	p2 := retPram(i+2, program, c.P2Immediate)
	p3 := retPram(i+3, program, true)

	program[p3] = p1 + p2
	return i + 4
}

func multiply(i int, c intermediatInstruct, program []int) int {
	p1 := retPram(i+1, program, c.P1Immediate)
	p2 := retPram(i+2, program, c.P2Immediate)
	p3 := retPram(i+3, program, true)

	program[p3] = p1 * p2
	return i + 4
}

func input(i int, c intermediatInstruct, program []int) int {
	// input store at i+1
	p1 := retPram(i+1, program, true)
	fmt.Print("-> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	inInt, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	program[p1] = inInt
	return i + 2
}
func output(i int, c intermediatInstruct, program []int) int {
	p1 := retPram(i+1, program, c.P1Immediate)
	out := p1
	oStream(out)
	return i + 2
}

func jumpIfTrue(i int, c intermediatInstruct, program []int) int {
	// jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value
	// from the second parameter. Otherwise, it does nothing.
	p1 := retPram(i+1, program, c.P1Immediate)

	if p1 != 0 {
		p2 := retPram(i+2, program, c.P2Immediate)
		return p2

	}
	return i + 3

}
func jumpIfFalse(i int, c intermediatInstruct, program []int) int {
	// jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value
	// from the second parameter. Otherwise, it does nothing.
	p1 := retPram(i+1, program, c.P1Immediate)

	if p1 == 0 {
		p2 := retPram(i+2, program, c.P2Immediate)
		return p2

	}
	return i + 3

}

func lessThan(i int, c intermediatInstruct, program []int) int {
	// less than: if the first parameter is less than the second parameter,
	// it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	p1 := retPram(i+1, program, c.P1Immediate)
	p2 := retPram(i+2, program, c.P2Immediate)
	p3 := retPram(i+3, program, true)

	if p1 < p2 {
		program[p3] = 1
	} else {
		program[p3] = 0
	}
	return i + 4
}

func equal(i int, c intermediatInstruct, program []int) int {
	// equals: if the first parameter is equal to the second parameter,
	// it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	p1 := retPram(i+1, program, c.P1Immediate)
	p2 := retPram(i+2, program, c.P2Immediate)
	p3 := retPram(i+3, program, true)

	if p1 == p2 {
		program[p3] = 1
	} else {
		program[p3] = 0
	}
	return i + 4
}

func retPram(i int, program []int, immediate bool) int {
	if immediate {
		return program[i]
	}
	return program[program[i]]

}

type intermediatInstruct struct {
	Code        int
	P1Immediate bool
	P2Immediate bool
	P3Immediate bool
}

func BreakUpOpCode(code int) intermediatInstruct {

	c := ((digit(code, 2) * 10) + digit(code, 1))
	p1 := digit(code, 3) == 1
	p2 := digit(code, 4) == 1
	p3 := digit(code, 5) == 1

	op := intermediatInstruct{
		Code:        c,
		P1Immediate: p1,
		P2Immediate: p2,
		P3Immediate: p3,
	}
	return op
}

// https://stackoverflow.com/questions/46753736/extract-digits-at-a-certain-position-in-a-number
func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
