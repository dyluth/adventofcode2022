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

type bag struct {
	first  string
	second string
}

func (b *bag) asRunes() []rune {
	return []rune(fmt.Sprintf("%s%s", b.first, b.second))
}

func part1() int {
	s := 0
	bags := fruitpicker.Load(NewBag, false)
	for _, b := range bags {
		dups := b.findDuplicates()
		for _, d := range dups {
			s += score(d)
		}
	}
	return s
}

func part2() int {
	s := 0
	bags := fruitpicker.Load(NewBag, false)

	// go through in 3s
	for i := 0; i < len(bags); i = i + 3 {
		fmt.Printf("batch %v\n", i)
		s += manageGroup(bags[i], bags[i+1], bags[i+2])
	}
	return s
}

func manageGroup(a, b, c bag) int {
	s := 0
	common := findCommon(a.asRunes(), b.asRunes(), c.asRunes())
	for _, d := range common {
		s += score(d)
	}
	return (s)
}

func findCommon(a, b, c []rune) []rune {
	common := make(map[rune]any)
	for _, r := range a {
		if strings.ContainsRune(string(b), r) && strings.ContainsRune(string(c), r) {
			common[r] = nil
		}
	}
	for _, r := range b {
		if strings.ContainsRune(string(a), r) && strings.ContainsRune(string(c), r) {
			common[r] = nil
		}
	}
	for _, r := range c {
		if strings.ContainsRune(string(b), r) && strings.ContainsRune(string(a), r) {
			common[r] = nil
		}
	}

	res := []rune{}
	for c := range common {
		res = append(res, c)
	}
	return res
}

func NewBag(s string) bag {

	ir := []rune(s)
	return bag{
		first:  string(ir[:len(ir)/2]),
		second: string(ir[len(ir)/2:]),
	}
}

func (b *bag) findDuplicates() []rune {
	a := make(map[rune]any)
	for _, r := range b.first {
		if strings.ContainsRune(b.second, r) {
			a[r] = nil
		}
	}
	res := []rune{}
	for c := range a {
		res = append(res, c)
	}
	return res
}

func score(letter rune) int {
	if int(letter) > 96 {
		return int(letter) - 96
	}
	return int(letter) - 38

}
