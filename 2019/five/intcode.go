package five

import (
	"fmt"
	"log"
	"math"
)

// ProcessProgram takes a intcode program and runs it to completion or error
func ProcessProgram(i int, program []int) error {

	opcode := program[i]

	if i >= len(program) {
		return fmt.Errorf("overflowed program")
	}

	c := BreakUpOpCode(opcode)
	switch c.Code {
	case 1:
		// Add [input1, input2, output]
		p1 := 0
		if c.P1Immediate {
			p1 = program[i+1]
		} else {
			p1 = program[program[i+1]]
		}

		p2 := 0
		if c.P2Immediate {
			p2 = program[i+2]
		} else {
			p2 = program[program[i+2]]
		}

		p3 := program[i+3]

		program[p3] = p1 + p2
		return ProcessProgram(i+4, program)
	case 2:
		// Multiply [input1, input2, output]
		// Add [input1, input2, output]
		p1 := 0
		if c.P1Immediate {
			p1 = program[i+1]
		} else {
			p1 = program[program[i+1]]
		}

		p2 := 0
		if c.P2Immediate {
			p2 = program[i+2]
		} else {
			p2 = program[program[i+2]]
		}

		p3 := program[i+3]
		program[p3] = p1 * p2
		return ProcessProgram(i+4, program)
	case 3:
		// input store at i+1
		p1 := program[i+1]
		program[p1] = 5
		return ProcessProgram(i+2, program)
	case 4:
		// output store at i+1
		p1 := 0
		if c.P1Immediate {
			p1 = i + 1
		} else {
			p1 = program[i+1]
		}
		log.Println(program[p1])
		return ProcessProgram(i+2, program)
	case 5:
		// jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value
		// from the second parameter. Otherwise, it does nothing.
		p1 := 0
		if c.P1Immediate {
			p1 = program[i+1]
		} else {
			p1 = program[program[i+1]]
		}

		p2 := 0
		if c.P2Immediate {
			p2 = program[i+2]
		} else {
			p2 = program[program[i+2]]
		}

		if p1 != 0 {
			return ProcessProgram(p2, program)
		}
		return ProcessProgram(i+3, program)
	case 6:
		//jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value
		// from the second parameter. Otherwise, it does nothing.
		p1 := 0
		if c.P1Immediate {
			p1 = program[i+1]
		} else {
			p1 = program[program[i+1]]
		}

		p2 := 0
		if c.P2Immediate {
			p2 = program[i+2]
		} else {
			p2 = program[program[i+2]]
		}

		if p1 == 0 {
			return ProcessProgram(p2, program)
		}
		return ProcessProgram(i+3, program)
	case 7:
		// less than: if the first parameter is less than the second parameter,
		// it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
		p1 := 0
		if c.P1Immediate {
			p1 = program[i+1]
		} else {
			p1 = program[program[i+1]]
		}

		p2 := 0
		if c.P2Immediate {
			p2 = program[i+2]
		} else {
			p2 = program[program[i+2]]
		}

		p3 := program[i+3]

		if p1 < p2 {
			program[p3] = 1
		} else {
			program[p3] = 0
		}
		return ProcessProgram(i+4, program)
	case 8:
		// equals: if the first parameter is equal to the second parameter,
		// it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
		p1 := 0
		if c.P1Immediate {
			p1 = program[i+1]
		} else {
			p1 = program[program[i+1]]
		}

		p2 := 0
		if c.P2Immediate {
			p2 = program[i+2]
		} else {
			p2 = program[program[i+2]]
		}

		p3 := program[i+3]

		if p1 == p2 {
			program[p3] = 1
		} else {
			program[p3] = 0
		}
		return ProcessProgram(i+4, program)
	case 99:
		return nil
	default:
		return fmt.Errorf("unknown opcode: %d", opcode)
	}

}

type opCode struct {
	Code        int
	P1Immediate bool
	P2Immediate bool
	P3Immediate bool
}

func BreakUpOpCode(code int) opCode {

	c := ((digit(code, 2) * 10) + digit(code, 1))
	p1 := digit(code, 3) == 1
	p2 := digit(code, 4) == 1
	p3 := digit(code, 5) == 1

	op := opCode{
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

// ProcessGravityAssist process the gravityAssistProgram using the given now and verb
func ProcessGravityAssist(noun, verb int) int {
	// I need a slice
	p := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 9, 19, 23, 1, 13, 23, 27, 1, 5, 27, 31, 2, 31, 6, 35, 1, 35, 5, 39, 1, 9, 39, 43, 1, 43, 5, 47, 1, 47, 5, 51, 2, 10, 51, 55, 1, 5, 55, 59, 1, 59, 5, 63, 2, 63, 9, 67, 1, 67, 5, 71, 2, 9, 71, 75, 1, 75, 5, 79, 1, 10, 79, 83, 1, 83, 10, 87, 1, 10, 87, 91, 1, 6, 91, 95, 2, 95, 6, 99, 2, 99, 9, 103, 1, 103, 6, 107, 1, 13, 107, 111, 1, 13, 111, 115, 2, 115, 9, 119, 1, 119, 6, 123, 2, 9, 123, 127, 1, 127, 5, 131, 1, 131, 5, 135, 1, 135, 5, 139, 2, 10, 139, 143, 2, 143, 10, 147, 1, 147, 5, 151, 1, 151, 2, 155, 1, 155, 13, 0, 99, 2, 14, 0, 0}
	// replace position 1 with the value 12 and replace position 2 with the value 2.
	p[1] = noun
	p[2] = verb
	err := ProcessProgram(0, p)
	if err != nil {
		log.Println(err)
	}
	return p[0]
}

// FindNounAndVerbForOutput finds a noun and verb for the gravity Assist program that will result in the wanted output
func FindNounAndVerbForOutput(output int, min, max int) (noun, verb int) {
	// Brut force
	for n := min; n < max; n++ {
		for v := min; v < max; v++ {
			o := ProcessGravityAssist(n, v)
			if o == output {
				return n, v
			}
		}
	}

	return -1, -1
}
