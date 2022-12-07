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
	return DoWindow(4)
}
func part2() int {
	return DoWindow(14)
}

func DoWindow(windowSize int) int {
	stream := []rune(fruitpicker.ReadInput()[0])
	for i := range stream {
		window := stream[i : i+windowSize]
		unique := true
		for j := range window {
			if strings.Contains(string(window[j+1:]), string(window[j])) {
				unique = false
				break
			}
			// all unique
		}
		if unique {
			return i + windowSize
		}
	}
	panic("dunno")

}
