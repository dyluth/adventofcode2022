package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dyluth/adventofcode2022/fruitpicker"
)

func main() {
	fmt.Printf("\nPART 1: " + part1() + "\n")
	fmt.Printf("\nPART 2: " + part2() + "\n")
}

func part1() string {
	input := fruitpicker.ReadGroupedInput()
	stack := NewStacks(input[0], input[1])

	for i := range stack.instructions {
		stack.state = stack.instructions[i].do(stack.state)
	}

	return stack.TopOfEach()
}

func part2() string {
	input := fruitpicker.ReadGroupedInput()
	stack := NewStacks(input[0], input[1])

	for i := range stack.instructions {
		stack.state = stack.instructions[i].doFancy(stack.state)
	}

	return stack.TopOfEach()
}

var stackRE = regexp.MustCompile(`.(.).\s`)
var cRE = regexp.MustCompile(`\[(.)\]`)

func NewStacks(diagram, instructions string) Stack {
	diagramSplit := strings.Split(diagram, "\n")

	for i := range diagramSplit {
		diagramSplit[i] = fmt.Sprintf("%v ", diagramSplit[i])
	}

	// start at the bottom
	// the bottom row is the row number - they just count up - we can ignore that, except to get the length

	matches := stackRE.FindAllString(diagramSplit[len(diagramSplit)-1], -1)
	state := make([][]string, len(matches)+1)

	for i := len(diagramSplit) - 2; i > -1; i-- {
		matches := stackRE.FindAllString(diagramSplit[i], -1)
		for j := range matches {
			bit := cRE.FindStringSubmatch(matches[j])
			if len(bit) > 0 {
				if state[j+1] == nil {
					state[j+1] = []string{}
				}
				state[j+1] = append(state[j+1], bit[1])
			}
		}
	}
	ins := LoadInstructions(strings.Split(instructions, "\n"))

	return Stack{state: state, instructions: ins}
}

func LoadInstructions(in []string) []Instruction {
	return fruitpicker.Map(in, ParseInstruction)
}

type Instruction struct {
	from  int
	to    int
	count int
}

var insRE = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

// move 1 from 2 to 1
func ParseInstruction(s string) Instruction {
	bits := insRE.FindStringSubmatch(s)
	val1, _ := strconv.Atoi(bits[1])
	val2, _ := strconv.Atoi(bits[2])
	val3, _ := strconv.Atoi(bits[3])
	return Instruction{
		from:  val2,
		to:    val3,
		count: val1,
	}
}

func (in *Instruction) doFancy(state [][]string) [][]string {

	from, end := state[in.from][:len(state[in.from])-(in.count)], state[in.from][len(state[in.from])-(in.count):]
	state[in.from] = from
	state[in.to] = append(state[in.to], end...)

	return state
}

func (in *Instruction) do(state [][]string) [][]string {

	for i := 0; i < in.count; i++ {
		from, end := state[in.from][:len(state[in.from])-1], state[in.from][len(state[in.from])-1]
		state[in.from] = from
		state[in.to] = append(state[in.to], end)
	}
	return state
}

type Stack struct {
	state        [][]string
	instructions []Instruction
}

func (s Stack) TopOfEach() string {
	out := []string{}
	for i := range s.state {
		if len(s.state[i]) > 0 {
			out = append(out, s.state[i][len(s.state[i])-1])
		}
	}
	return strings.Join(out, "")
}
