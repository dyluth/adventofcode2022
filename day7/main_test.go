package main

import (
	"strings"
	"testing"

	"github.com/dyluth/adventofcode2022/fruitpicker"
	"github.com/stretchr/testify/require"
)

var all = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func Test_structure2(t *testing.T) {
	fruitpicker.LoadStringOverride = all
	p := part1()
	require.Equal(t, 95437, p)
	require.True(t, false)
}

func Test_structure(t *testing.T) {
	root := structure(strings.Split(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.da`, "\n"))
	root.print(0)
	require.Len(t, root.dirs, 0)
	require.Len(t, root.files, 2)
	f1, ok := root.files["b.txt"]
	require.True(t, ok)
	require.Equal(t, 14848514, f1.size)
	f2, ok := root.files["c.da"]
	require.True(t, ok)
	require.Equal(t, 8504156, f2.size)
}

var all2 = `$ cd /
$ ls
dir a
dir c
$ cd a
$ ls
dir b
2 b
$ cd b
$ ls
5 i
$ cd ..
$ cd ..
$ cd c
$ ls
1 j`

// /(8)
//
//	a(7)
//		2
//		b (5)
//			5
//	c(1)
//		1
//
// 8 + 7 + 5  + 1 = 21
func Test_structure3(t *testing.T) {
	fruitpicker.LoadStringOverride = all2
	p := part1()
	require.Equal(t, 21, p)
}
