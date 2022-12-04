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

func part2() int {
	elves := fruitpicker.Load(func(s string) int {
		return fruitpicker.SumList(fruitpicker.ConvertStringsToInts(strings.Split(s, "\n")))
	}, true)

	largest, index := fruitpicker.Largest(elves)
	sum := largest
	// remove that elf
	elves = append(elves[0:index], elves[index+1:]...)

	largest, index = fruitpicker.Largest(elves)
	sum += largest
	// remove that elf
	elves = append(elves[0:index], elves[index+1:]...)

	largest, _ = fruitpicker.Largest(elves)
	sum += largest

	return sum
}

func part1() int {
	elves := fruitpicker.Load(func(s string) int {
		return fruitpicker.SumList(fruitpicker.ConvertStringsToInts(strings.Split(s, "\n")))
	}, true)

	largest, _ := fruitpicker.Largest(elves)
	return largest
}
