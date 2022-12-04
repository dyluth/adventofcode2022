package main

import (
	"testing"

	"github.com/dyluth/adventofcode2022/fruitpicker"
	"github.com/stretchr/testify/require"
)

func Test_part1(t *testing.T) {
	fruitpicker.LoadStringOverride = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	require.Equal(t, 2, part1())
}

func Test_part2(t *testing.T) {
	fruitpicker.LoadStringOverride = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	require.Equal(t, 4, part2())
}

func TestPair_IsContained(t *testing.T) {
	require.False(t, NewPair("2-3,4-5").IsContained())
	require.True(t, NewPair("2-5,4-5").IsContained())
	require.True(t, NewPair("4-7,4-5").IsContained())
	require.True(t, NewPair("4-5,4-5").IsContained())
}
