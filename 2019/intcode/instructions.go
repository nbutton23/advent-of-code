package intcode

/*
add

OPCode: 1
add P1 P2 and stores in P3
*/
func (p *Process) add(i int, c instruction) int {
	p1, p2, p3 := p.getDefaultParamSet(i, c)
	p.setValue(p3, p1+p2)

	return i + 4
}

/*
multiply

OPCode: 2
multiplies P1 P2 and stores in P3
*/
func (p *Process) multiply(i int, c instruction) int {
	p1, p2, p3 := p.getDefaultParamSet(i, c)
	p.setValue(p3, p1*p2)

	return i + 4
}

/*
input

OPCode: 3
Stores the input from the inStream to P1
*/
func (p *Process) input(i int, c instruction) int {
	p1 := p.indexForPram(i+1, p.program, c.P1Mode)

	in := p.inStream()
	p.setValue(p1, in)
	return i + 2
}

/*
output

OPCode: 4
Outputs P1 to the oStream
*/
func (p *Process) output(i int, c instruction) int {
	p1 := p.valueForPram(i+1, p.program, c.P1Mode)
	out := p1
	p.outputs = append(p.outputs, out)
	p.oStream(out)
	return i + 2
}

/*
jumpIfTrue

OPCode: 5
Jump to P2 if P1 != 0
*/
func (p *Process) jumpIfTrue(i int, c instruction) int {
	p1, p2, _ := p.getDefaultParamSet(i, c)

	if p1 != 0 {
		return p2

	}
	return i + 3
}

/*
jumpIfFalse

OPCode: 6
Jump to P2 if P1 == 0
*/
func (p *Process) jumpIfFalse(i int, c instruction) int {
	p1, p2, _ := p.getDefaultParamSet(i, c)

	if p1 == 0 {
		return p2

	}
	return i + 3
}

/*
lessThan

OPCode: 7
if P1 < P2  store 1 in P3 else store 0 in P3
*/
func (p *Process) lessThan(i int, c instruction) int {
	p1, p2, p3 := p.getDefaultParamSet(i, c)

	if p1 < p2 {
		p.setValue(p3, 1)
	} else {
		p.setValue(p3, 0)
	}
	return i + 4
}

/*
equal

OPCode: 8
if P1 == P2  store 1 in P3 else store 0 in P3
*/
func (p *Process) equal(i int, c instruction) int {
	p1, p2, p3 := p.getDefaultParamSet(i, c)

	if p1 == p2 {
		p.setValue(p3, 1)
	} else {
		p.setValue(p3, 0)
	}
	return i + 4
}

/*
adjustRelativeBase

OPCode: 9
increases (or decreases, if the value is negative) the relitive base of the process by P1
*/
func (p *Process) adjustRelativeBase(i int, c instruction) int {
	p1 := p.valueForPram(i+1, p.program, c.P1Mode)

	p.relativeBaseOffset += p1
	return i + 2
}

/*
halt

OPCode: 99
Halt program excution
*/
func (p *Process) halt(i int, c instruction) int {
	p.IsHalted = true
	return -1
}
