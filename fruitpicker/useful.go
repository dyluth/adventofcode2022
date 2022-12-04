package fruitpicker

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func Reduce[T any](t []T, f func(T) int) int {
	total := 0
	for _, v := range t {
		total += f(v)
	}
	return total
}

// add up a list
func SumList[T Number](in []T) (total T) {
	for i := range in {
		total = total + in[i]
	}
	return total
}

type ValueAble interface {
	Value() int
}

func Largest[T ValueAble](in []T) (most T) {
	val := 0
	for _, v := range in {
		if v.Value() > val {
			val = v.Value()
			most = v
		}
	}
	return most
}

func Select[T any](in []T, f func(i T) bool) (out []T) {
	for i := range in {
		if f(in[i]) {
			out = append(out, in[i])
		}
	}
	return
}
