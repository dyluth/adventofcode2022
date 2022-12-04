package fruitpicker

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func Map[T, S any](t []T, f func(T) S) []S {
	r := make([]S, len(t))
	for i := range t {
		r[i] = f(t[i])
	}
	return r
}

func Reduce[T, S any](t []T, accumulator S, f func(T, S, int) S) S {
	for i, v := range t {
		accumulator = f(v, accumulator, i)
	}
	return accumulator
}

// add up a list
func SumList[T Number](in []T) (total T) {
	return Reduce(in, 0, func(cur, acc T, i int) T {
		return cur + acc
	})
}

func Largest[T Number](in []T) (T, int) {
	if len(in) == 0 {
		panic("wtf, 0 length slice?")
	}
	var index int
	largest := Reduce(in, in[0], func(cur, acc T, i int) T {
		if cur > acc {
			index = i
			return cur
		}
		return acc
	})
	return largest, index
}

func Smallest[T Number](in []T) (T, int) {
	if len(in) == 0 {
		panic("wtf, 0 length slice?")
	}
	var index int
	largest := Reduce(in, in[0], func(cur, acc T, i int) T {
		if cur < acc {
			index = i
			return cur
		}
		return acc
	})
	return largest, index
}

func Select[T any](in []T, f func(i T) bool) (out []T) {
	for i := range in {
		if f(in[i]) {
			out = append(out, in[i])
		}
	}
	return
}
