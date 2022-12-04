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

func Reduce[T, S any](t []T, accumulator S, f func(T, S) S) S {
	for _, v := range t {
		accumulator = f(v, accumulator)
	}
	return accumulator
}

// add up a list
func SumList[T Number](in []T) (total T) {
	sum := func(a T, b T) T {
		return a + b
	}
	return Reduce(in, 0, sum)
}

type ValueAble[T Number] interface {
	Value() T
}

func Largest[T Number, S ValueAble[T]](in []S) (most S) {
	if len(in) == 0 {
		panic("wtf, 0 length slice?")
	}
	return Reduce(in, in[0], func(cur, acc S) S {
		if cur.Value() > acc.Value() {
			return cur
		}
		return acc
	})
}

func Select[T any](in []T, f func(i T) bool) (out []T) {
	for i := range in {
		if f(in[i]) {
			out = append(out, in[i])
		}
	}
	return
}
