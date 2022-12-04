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
	pairs := fruitpicker.Load(NewPair, false)
	return fruitpicker.Reduce(pairs, 0, func(p Pair, acc, i int) int {
		if p.IsContained() {
			return acc + 1
		}
		return acc
	})
}

func part2() int {
	pairs := fruitpicker.Load(NewPair, false)
	return fruitpicker.Reduce(pairs, 0, func(p Pair, acc, i int) int {
		if p.Overlap() {
			return acc + 1
		}
		return acc
	})
}

type Pair struct {
	p1a, p1b, p2a, p2b int
}

func (p Pair) IsContained() bool {
	if p.p1a >= p.p2a && p.p1b <= p.p2b {
		return true
	}
	if p.p1a <= p.p2a && p.p1b >= p.p2b {
		return true
	}
	return false
}

// 2-5, 4-7
// 4-7, 2,5
func (p Pair) Overlap() bool {
	// is p1a in the range of p2
	if p.p1a >= p.p2a && p.p1a <= p.p2b {
		return true
	}
	// is p1b in the range of p2
	if p.p1b >= p.p2a && p.p1b <= p.p2b {
		return true
	}
	if p.IsContained() {
		return true
	}
	return false
}

// expects format:  2-4,6-8
func NewPair(s string) Pair {
	ps := strings.Split(s, ",")
	p1 := fruitpicker.ConvertStringsToInts(strings.Split(ps[0], "-"))
	p2 := fruitpicker.ConvertStringsToInts(strings.Split(ps[1], "-"))
	return Pair{
		p1a: p1[0],
		p1b: p1[1],
		p2a: p2[0],
		p2b: p2[1],
	}
}
