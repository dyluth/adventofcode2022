package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dyluth/adventofcode2022/fruitpicker"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part2() int {
	elves := fruitpicker.Load(func(s string) Elf {
		return Elf{
			carrying: fruitpicker.ConvertStringsToInts(strings.Split(s, "\n")),
		}
	}, true)

	largest := fruitpicker.Largest(elves)
	sum := largest.Value()
	// remove that elf
	elves = fruitpicker.Select(elves, func(e Elf) bool {
		return !reflect.DeepEqual(largest, e)
	})

	largest = fruitpicker.Largest(elves)
	sum += largest.Value()
	elves = fruitpicker.Select(elves, func(e Elf) bool {
		return !reflect.DeepEqual(largest, e)
	})
	largest = fruitpicker.Largest(elves)
	sum += largest.Value()

	return sum

}

func part1() int {
	elves := fruitpicker.Load(func(s string) Elf {
		return Elf{
			carrying: fruitpicker.ConvertStringsToInts(strings.Split(s, "\n")),
		}
	}, true)

	e := fruitpicker.Largest(elves)
	return e.Value()
}

type Elf struct {
	carrying []int
}

func (e Elf) Value() int {
	return fruitpicker.SumList(e.carrying)
}

//func findH
