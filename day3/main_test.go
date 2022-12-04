package main

import (
	"testing"

	"github.com/dyluth/adventofcode2022/fruitpicker"
	"github.com/stretchr/testify/require"
)

func Test_bag_findDuplicates(t *testing.T) {
	b := NewBag("vJrwpWtwJgWrhcsFMMfFFhFp")
	dups := b.findDuplicates()
	require.Len(t, dups, 1)
	require.Equal(t, string(dups[0]), "p")
}

func Test_NewBag_findDuplicates(t *testing.T) {
	b := NewBag("aaabbb")
	require.Equal(t, "aaa", b.first)
	require.Equal(t, "bbb", b.second)
}

func Test_score(t *testing.T) {
	require.Equal(t, 1, score(rune('a')))
	require.Equal(t, 26, score(rune('z')))
	require.Equal(t, 27, score(rune('A')))
	require.Equal(t, 52, score(rune('Z')))
}

func Test_part1(t *testing.T) {
	fruitpicker.LoadStringOverride = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	require.Equal(t, 157, part1())
}

func Test_part2(t *testing.T) {
	fruitpicker.LoadStringOverride = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	require.Equal(t, 70, part2())
}

func Test_findCommon(t *testing.T) {
	r := findCommon([]rune("vJrwpWtwJgWrhcsFMMfFFhFp"), []rune("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"), []rune("PmmdzqPrVvPwwTWBwg"))
	require.Len(t, r, 1)
	require.Equal(t, "r", string(r))
}

func Test_manageGroup(t *testing.T) {
	s := manageGroup(NewBag("vJrwpWtwJgWrhcsFMMfFFhFp"), NewBag("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"), NewBag("PmmdzqPrVvPwwTWBwg"))
	require.Equal(t, 18, s)
}
