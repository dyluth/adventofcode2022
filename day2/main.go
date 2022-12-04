package main

import (
	"fmt"
	"strings"

	"github.com/dyluth/adventofcode2022/fruitpicker"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	lines := fruitpicker.Load(func(s string) RPS {
		bits := strings.Fields(s)
		r := RPS{
			them: bits[0],
			us:   bits[1],
		}
		return r
	}, false)

	score := 0
	for _, r := range lines {
		score += r.getScore()
	}
	return score
}

type RPS struct {
	them string
	us   string
}

func (r *RPS) getScore() int {
	score := 0
	offset1 := 0
	offset2 := 0
	switch r.us {
	case "X":
		score += 1
		offset1 = 0
	case "Y":
		score += 2
		offset1 = 1
	case "Z":
		score += 3
		offset1 = 2
	}

	switch r.them {
	case "A":
		offset2 = 0

	case "B":
		offset2 = 1

	case "C":
		offset2 = 2
	}

	result := offset1 - offset2
	if result == 0 { // draw
		score += 3
	} else if result == 1 || result == -2 { // won
		score += 6
	} else {
		score += 0
	}

	return score
}

func part2() int {
	lines := fruitpicker.Load(func(s string) RPS {
		bits := strings.Fields(s)
		r := RPS{
			them: bits[0],
			us:   bits[1],
		}
		return r
	}, false)

	score := 0
	for _, r := range lines {
		score += r.part2()
	}
	return score
}

func (r *RPS) part2() int {
	score := 0
	offset1 := 0
	offset2 := 0

	switch r.them {
	case "A":
		offset2 = 0

	case "B":
		offset2 = 1

	case "C":
		offset2 = 2
	}

	switch r.us { // what the resuls should be
	case "X": // lose
		offset1 = offset2 - 1
		score += 0
	case "Y": // draw
		offset1 = offset2
		score += 3
	case "Z": // win
		offset1 = offset2 + 1
		score += 6
	}

	offset1 = (offset1 + 3) % 3

	if offset1 == 0 { // Rock
		score += 1
	} else if offset1 == 1 { // Paper
		score += 2
	} else if offset1 == 2 { // Scissors
		score += 3
	}

	return score

}
