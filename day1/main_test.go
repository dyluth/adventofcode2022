package main

import (
	"testing"

	"github.com/dyluth/adventofcode2022/fruitpicker"
	"github.com/stretchr/testify/require"
)

func Test_part1(t *testing.T) {
	fruitpicker.LoadStringOverride = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	require.Equal(t, 24000, part1())
}

func Test_part2(t *testing.T) {
	fruitpicker.LoadStringOverride = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	require.Equal(t, 45000, part2())
}
