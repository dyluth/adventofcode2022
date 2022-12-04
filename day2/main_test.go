package main

import (
	"testing"

	"github.com/dyluth/adventofcode2022/fruitpicker"
	"github.com/stretchr/testify/require"
)

func Test_part1(t *testing.T) {
	fruitpicker.LoadStringOverride = `A Y
B X
C Z`
	require.Equal(t, 15, part1())
}

func TestRPS_part2(t *testing.T) {
	fruitpicker.LoadStringOverride = `A Y
B X
C Z`
	require.Equal(t, 12, part2())
}
