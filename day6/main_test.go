package main

import (
	"testing"

	"github.com/dyluth/adventofcode2022/fruitpicker"
	"github.com/stretchr/testify/require"
)

func Test_part1(t *testing.T) {
	fruitpicker.LoadStringOverride = "bvwbjplbgvbhsrlpgdmjqwftvncz"
	require.Equal(t, 5, part1())
	fruitpicker.LoadStringOverride = "nppdvjthqldpwncqszvftbrmjlhg"
	require.Equal(t, 6, part1())
	fruitpicker.LoadStringOverride = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	require.Equal(t, 10, part1())
	fruitpicker.LoadStringOverride = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	require.Equal(t, 11, part1())

}

func Test_part2(t *testing.T) {
	fruitpicker.LoadStringOverride = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	require.Equal(t, 26, part2())
}
