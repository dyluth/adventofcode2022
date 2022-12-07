package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dyluth/adventofcode2022/fruitpicker"
	"github.com/stretchr/testify/require"
)

func TestNewStacks(t *testing.T) {
	fruitpicker.LoadStringOverride = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	part1()
	require.True(t, false)
}

func TestLoadInstructions(t *testing.T) {
	s := `move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	ins := LoadInstructions(strings.Split(s, "\n"))
	fmt.Printf("INS: %+v", ins)
	require.Len(t, ins, 4)
	require.Equal(t, Instruction{count: 1, from: 2, to: 1}, ins[0])
}

func TestInstruction_do(t *testing.T) {
	ins := Instruction{from: 0, to: 1, count: 2}
	state := [][]string{
		{
			"a", "b",
		},
		{},
		{"c", "d"},
	}
	s2 := ins.do(state)
	fmt.Printf("STATE:\n%+v\n", state)
	fmt.Printf("STATE2:\n%+v\n", s2)
	require.Equal(t, 3, len(s2))
	require.Equal(t, 0, len(s2[0]))
	require.Equal(t, 2, len(s2[1]))
	require.Equal(t, "b", s2[1][0])
	require.Equal(t, "a", s2[1][1])

	s := Stack{state: state}
	top := s.TopOfEach()
	require.Equal(t, "ad", top)
}
