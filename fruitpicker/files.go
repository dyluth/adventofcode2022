package fruitpicker

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

var LoadStringOverride = ""

func Load[T any](convert func(string) T, grouped bool) []T {
	out := []T{}
	loaded := []string{}
	if grouped {
		loaded = ReadGroupedInput()
	} else {
		loaded = ReadInput()
	}
	for i := range loaded {
		out = append(out, convert(loaded[i]))
	}
	return out
}

// ReadInput returns lines of input separated by new lines
func ReadInput() []string {
	data := read()
	return strings.Split(data, "\n")
}

// ConvertStringsToInts - converts input to ints if appropriate
func ConvertStringsToInts(s []string) []int {
	ints := make([]int, len(s))
	for i, v := range s {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		ints[i] = val
	}
	return ints
}

// ReadGroupedInput returns groups of input separated by empty lines
func ReadGroupedInput() []string {
	data := read()
	re := regexp.MustCompile(`\n\s*\n`)
	passList := re.Split(data, -1)
	return passList
}

func read() string {
	if LoadStringOverride != "" {
		return LoadStringOverride
	}
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return string(dat)
}
