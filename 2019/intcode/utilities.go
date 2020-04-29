package intcode

func (p *Process) getValue(index int) int {
	if index < 0 {
		panic("Negative index!")
	}

	if index >= len(p.program) {
		p.expandMemory(index + 1)
	}
	return p.program[index]
}

func (p *Process) setValue(index int, value int) {
	if index < 0 {
		panic("Negative index!")
	}

	if index >= len(p.program) {
		p.expandMemory(index + 1)
	}

	p.program[index] = value
}

func (p *Process) expandMemory(min int) {
	expand := len(p.program) * 2
	if expand < min {
		expand = min
	}
	// double the length
	// TODO: We can probably do better than this
	tempP := make([]int, expand)
	copy(tempP, p.program)
	p.program = tempP
}

// Returns the Value a value.
func (p *Process) valueForPram(i int, program []int, mode accessMode) int {
	index := p.indexForPram(i, program, mode)
	return p.getValue(index)
}

// Returns the index
func (p *Process) indexForPram(i int, program []int, mode accessMode) int {
	switch mode {
	case immediateMode:
		return i
	case positionMode:
		return p.getValue(i)
	case relativeMode:
		pos := p.relativeBaseOffset + p.getValue(i)
		return pos
	default:
		panic("Received unknown mode")
	}
}

/*
getDefaultParamSet

Returns p1 and p2 as valueFrom and P3 as indexFrom with the correct accessMode
*/
func (p *Process) getDefaultParamSet(i int, c instruction) (p1, p2, p3 int) {
	p1 = p.valueForPram(i+1, p.program, c.P1Mode)
	p2 = p.valueForPram(i+2, p.program, c.P2Mode)
	p3 = p.indexForPram(i+3, p.program, c.P3Mode)

	return p1, p2, p3
}
