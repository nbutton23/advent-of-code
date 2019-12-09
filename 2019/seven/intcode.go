package seven

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
)

var reader = bufio.NewReader(os.Stdin)
var defaultNextInt = func() int {
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

func MaxAmp(p []int, minPhase, maxPhase int) (int, error) {

	inputPhase := make([]int, 0)
	for i := minPhase; i <= maxPhase; i++ {
		inputPhase = append(inputPhase, i)
	}

	perms := permutations(inputPhase)
	max := -1

	for _, c := range perms {
		o, err := ProcessAmps(p, c, 0)
		if err != nil {
			return -1, err
		}

		if o > max {
			max = o
		}
	}

	return max, nil
}

func MaxAmpLoop(p []int, minPhase, maxPhase int) int {
	inputPhase := make([]int, 0)
	for i := minPhase; i <= maxPhase; i++ {
		inputPhase = append(inputPhase, i)
	}

	perms := permutations(inputPhase)
	max := -1

	for _, c := range perms {
		o := ProcessAmpLoop(p, c, 0)

		if o > max {
			max = o
		}
	}

	return max
}

type Process struct {
	tag        string
	inputChan  chan int
	outputChan chan int
	p          []int
}

func getTagForInt(i int) string {
	switch i {
	case 0:
		return "A"
	case 1:
		return "B"
	case 2:
		return "C"
	case 3:
		return "D"
	case 4:
		return "E"
	default:
		return "UNEXPECTED"

	}
}
func ProcessAmpLoop(p []int, phaseOrder []int, initInput int) int {

	amps := make([]Process, len(phaseOrder))
	output := make(chan int, 10)
	input := make(chan int, 10)
	initInputChan := input
	for i, proc := range phaseOrder {
		pCopy := make([]int, len(p))
		copy(pCopy, p)
		input <- proc
		amps[i] = Process{
			inputChan:  input,
			outputChan: output,
			p:          pCopy,
			tag:        getTagForInt(i),
		}

		input = output
		output = make(chan int, 10)
	}

	lock := sync.Mutex{}
	amps[len(amps)-1].outputChan = initInputChan

	wg := &sync.WaitGroup{}
	rets := make(map[string]int)
	for _, proc := range amps {
		wg.Add(1)
		go func(amp Process, wg *sync.WaitGroup) {
			l, err := ProcessProgramForImmediateOutputWithImmediateInput(amp.p, amp.inputChan, amp.outputChan, amp.tag)
			if err != nil {
				log.Println(err)
			}
			lock.Lock()
			rets[amp.tag] = l
			lock.Unlock()
			// log.Printf("[%s] - Exited\n", amp.tag)
			wg.Done()
		}(proc, wg)
	}
	initInputChan <- initInput
	wg.Wait()

	// o := <-output

	return rets["E"]

}
func ProcessAmps(p []int, phaseOrder []int, initInput int) (int, error) {

	pCopy := make([]int, len(p))

	outPut := initInput
	for _, phase := range phaseOrder {
		copy(pCopy, p)
		o, err := ProcessProgramForOutputWithInput(pCopy, []int{phase, outPut})
		if err != nil {
			return -1, err
		}

		outPut = o
	}

	return outPut, nil
}

func ProcessProgramForOutputWithInput(p []int, input []int) (int, error) {
	content := ""
	for _, i := range input {
		content = fmt.Sprintf("%s%d\n", content, i)
	}

	inputFile, err := ioutil.TempFile("temp", "input")
	if err != nil {
		return -1, err
	}
	if _, err := inputFile.Write([]byte(content)); err != nil {
		return -1, err
	}
	if _, err := inputFile.Seek(0, 0); err != nil {
		return -1, err
	}
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = inputFile
	defer os.Remove(inputFile.Name()) // clean up

	return ProcessProgramForOutput(p, nil)
}

func ProcessProgramForImmediateOutputWithImmediateInput(p []int, input chan int, output chan<- int, tag string) (int, error) {

	nextInt := func() int {
		// fmt.Printf("[%s] - Waiting for input\n", tag)
		o := <-input
		// fmt.Printf("[%s] - Got  input: %d\n", tag, o)
		return o
	}
	return ProcessProgramForImmediateOutput(p, output, nextInt, tag)
}
func ProcessProgramForImmediateOutput(p []int, output chan<- int, input func() int, tag string) (int, error) {
	oStream := func(a ...interface{}) (n int, err error) {
		o, ok := a[0].(int)
		if !ok {
			panic("expected int")
		}

		// fmt.Printf("[%s] - Writing output\n", tag)
		output <- o
		// fmt.Printf("[%s] - Output Wrote: %d\n", tag, o)

		return 0, nil
	}

	program := NewP(input, oStream)
	l, err := program.ProcessProgram(0, p)
	close(output)
	return l, err
}

func ProcessProgramForOutput(p []int, input func() int) (int, error) {
	output := make([]int, 0)
	oChan := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for o := range oChan {
			output = append(output, o)
		}
		wg.Done()
	}()
	l, err := ProcessProgramForImmediateOutput(p, oChan, nil, "")
	wg.Wait()
	return l, err
}

type P struct {
	nextInt func() int
	oStream func(a ...interface{}) (n int, err error)

	cmdsRan        bytes.Buffer
	instructionSet map[int]func(i int, c intermediatInstruct, program []int) int
	decomp         bool

	outputs []int
}

func NewP(nextInt func() int, oStream func(a ...interface{}) (n int, err error)) *P {

	if nextInt == nil {
		nextInt = defaultNextInt
	}

	if oStream == nil {
		oStream = fmt.Println
	}

	p := &P{
		nextInt: nextInt,
		decomp:  true,
		oStream: oStream,
		outputs: make([]int, 0),
	}
	p.instructionSet = map[int]func(i int, c intermediatInstruct, program []int) int{
		1:  p.add,
		2:  p.multiply,
		3:  p.input,
		4:  p.output,
		5:  p.jumpIfTrue,
		6:  p.jumpIfFalse,
		7:  p.lessThan,
		8:  p.equal,
		99: p.halt,
	}

	return p
}

// ProcessProgram takes a intcode program and runs it to completion or error
func (p *P) ProcessProgram(i int, program []int) (int, error) {
	p.cmdsRan = bytes.Buffer{}
	for i < len(program) && i >= 0 {
		c := BreakUpOpCode(program[i])
		if cmd, ok := p.instructionSet[c.Code]; ok {
			i = cmd(i, c, program)
		} else {
			return 0, fmt.Errorf("unknown instruction: %v", c)
		}
	}

	tmpfile, _ := ioutil.TempFile("temp", "program-output")
	p.cmdsRan.WriteString(fmt.Sprintf("%v", p.outputs))
	tmpfile.Write(p.cmdsRan.Bytes())
	tmpfile.Close()
	lastOuput := -1
	if len(p.outputs) > 0 {
		lastOuput = p.outputs[len(p.outputs)-1]
	}
	return lastOuput, nil
}

func (p *P) halt(i int, c intermediatInstruct, program []int) int {
	if p.decomp {
		p.cmdsRan.WriteString("HALT\n")
	}
	return -1
}

func (p *P) add(i int, c intermediatInstruct, program []int) int {
	// Add [input1, input2, output]
	p1 := retPram(i+1, program, c.P1Immediate)
	p2 := retPram(i+2, program, c.P2Immediate)
	p3 := retPram(i+3, program, true)

	program[p3] = p1 + p2

	if p.decomp {
		pr1 := retPramString(i+1, c.P1Immediate)
		pr2 := retPramString(i+2, c.P2Immediate)
		pr3 := retPramString(i+3, true)

		p.cmdsRan.WriteString(fmt.Sprintf("ADD %s %s %s\n", pr1, pr2, pr3))
	}
	return i + 4
}

func (p *P) multiply(i int, c intermediatInstruct, program []int) int {
	p1 := retPram(i+1, program, c.P1Immediate)
	p2 := retPram(i+2, program, c.P2Immediate)
	p3 := retPram(i+3, program, true)

	program[p3] = p1 * p2

	if p.decomp {
		pr1 := retPramString(i+1, c.P1Immediate)
		pr2 := retPramString(i+2, c.P2Immediate)
		pr3 := retPramString(i+3, true)

		p.cmdsRan.WriteString(fmt.Sprintf("MUL %s %s %s\n", pr1, pr2, pr3))
	}

	return i + 4
}

func (p *P) input(i int, c intermediatInstruct, program []int) int {
	// input store at i+1
	p1 := retPram(i+1, program, true)

	program[p1] = p.nextInt()

	if p.decomp {
		pr1 := retPramString(i+1, true)

		p.cmdsRan.WriteString(fmt.Sprintf("IN %s\n", pr1))
	}

	return i + 2
}
func (p *P) output(i int, c intermediatInstruct, program []int) int {
	p1 := retPram(i+1, program, c.P1Immediate)
	out := p1
	p.outputs = append(p.outputs, out)
	p.oStream(out)
	if p.decomp {
		pr1 := retPramString(i+1, c.P1Immediate)

		p.cmdsRan.WriteString(fmt.Sprintf("OUT %s\n", pr1))
	}
	return i + 2
}

func (p *P) jumpIfTrue(i int, c intermediatInstruct, program []int) int {
	// jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value
	// from the second parameter. Otherwise, it does nothing.
	p1 := retPram(i+1, program, c.P1Immediate)

	if p1 != 0 {
		p2 := retPram(i+2, program, c.P2Immediate)
		return p2

	}
	if p.decomp {
		pr1 := retPramString(i+1, c.P1Immediate)
		pr2 := retPramString(i+2, c.P2Immediate)

		p.cmdsRan.WriteString(fmt.Sprintf("JEQ %s %s\n", pr1, pr2))
	}
	return i + 3

}
func (p *P) jumpIfFalse(i int, c intermediatInstruct, program []int) int {
	// jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value
	// from the second parameter. Otherwise, it does nothing.
	p1 := retPram(i+1, program, c.P1Immediate)

	if p1 == 0 {
		p2 := retPram(i+2, program, c.P2Immediate)
		return p2

	}
	if p.decomp {
		pr1 := retPramString(i+1, c.P1Immediate)
		pr2 := retPramString(i+2, c.P2Immediate)

		p.cmdsRan.WriteString(fmt.Sprintf("JNE %s %s\n", pr1, pr2))
	}
	return i + 3

}

func (p *P) lessThan(i int, c intermediatInstruct, program []int) int {
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

	if p.decomp {
		pr1 := retPramString(i+1, c.P1Immediate)
		pr2 := retPramString(i+2, c.P2Immediate)
		pr3 := retPramString(i+3, true)

		p.cmdsRan.WriteString(fmt.Sprintf("LES %s %s %s\n", pr1, pr2, pr3))
	}

	return i + 4
}

func (p *P) equal(i int, c intermediatInstruct, program []int) int {
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

	if p.decomp {
		pr1 := retPramString(i+1, c.P1Immediate)
		pr2 := retPramString(i+2, c.P2Immediate)
		pr3 := retPramString(i+3, true)

		p.cmdsRan.WriteString(fmt.Sprintf("EQL %s %s %s\n", pr1, pr2, pr3))
	}
	return i + 4
}

func retPram(i int, program []int, immediate bool) int {
	if immediate {
		return program[i]
	}
	return program[program[i]]

}

func retPramString(i int, immediate bool) string {
	if immediate {
		return fmt.Sprintf("*%d", i)
	}
	return fmt.Sprintf("%d", i)
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

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := int(0); i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, int(len(arr)))
	return res
}
