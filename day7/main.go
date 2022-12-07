package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dyluth/adventofcode2022/fruitpicker"
)

var debug = true

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	root := structure(fruitpicker.ReadInput())
	sizes := root.calculateSizes()
	total := 0
	for k := range sizes {
		if sizes[k] <= 100000 {
			total += sizes[k]
		}
	}
	fmt.Println(total)
	return total
}

func part2() int {
	root := structure(fruitpicker.ReadInput())
	consumed := root.calculateSize()
	free := 70000000 - consumed
	minimalToDelete := 30000000 - free

	// 70000000
	// need 30000000

	sizes := root.calculateSizes()
	smallest := consumed
	fmt.Printf("consumed: %v\n free: %v\nneed to free up %v\n", consumed, free, minimalToDelete)
	for k := range sizes {
		if sizes[k] > minimalToDelete && sizes[k] < smallest {
			smallest = sizes[k]
		}
	}
	return smallest
}

type dir struct {
	name   string
	parent *dir
	dirs   map[string]dir
	files  map[string]file
}

func (d *dir) print(indent int) {
	fmt.Printf("%vDIR: %v\n", strings.Repeat("  ", indent), d.name)
	for _, f := range d.files {
		fmt.Printf("%v%v - %v\n", strings.Repeat("  ", indent), f.name, f.size)
	}
	for _, v := range d.dirs {
		v.print(indent + 1)
	}
}

func (d *dir) calculateSize() int {
	s := 0
	for _, c := range d.dirs {
		s += c.calculateSize()
	}
	for _, f := range d.files {
		s += f.size
	}
	return s
}

func (d *dir) calculateSizes() []int {
	s := d.calculateSize()
	sizes := []int{}
	for _, c := range d.dirs {
		sizes = append(sizes, c.calculateSizes()...)
	}
	sizes = append(sizes, s)
	return sizes
}

type file struct {
	name string
	size int
}

// eg 8033020 d.log
func NewFile(in string) file {
	f := strings.Fields(in)
	s, err := strconv.Atoi(f[0])
	if err != nil {
		panic(err)
	}
	return file{
		name: f[1],
		size: s,
	}
}

// $ cd a
func NewDir(in string, parent *dir) *dir {
	f := strings.Fields(in)
	d := dir{
		name:   f[2],
		parent: parent,
		dirs:   make(map[string]dir),
		files:  make(map[string]file),
	}
	if parent != nil {
		_, ok := parent.dirs[d.name]
		if ok {
			panic(fmt.Sprintf("already exists! %v\n", d.name))
		}

		parent.dirs[d.name] = d
	}
	return &d
}

func structure(lines []string) dir {
	root := NewDir("$ cd /", nil)
	currentDir := root
	for i := 1; i < len(lines); i++ {
		// new command
		f := strings.Fields(lines[i])
		if f[0] == "$" {
			switch f[1] {
			case "cd":
				if f[2] == ".." {
					// go back up to parent
					currentDir = currentDir.parent
				} else if f[2] == "/" {
					currentDir = root
				} else {
					// create a new directory
					currentDir = NewDir(lines[i], currentDir)
				}
			case "ls":
				// keep peeking forwards till we find a $
				// yes we are ALSO increasing i here...
				for ; i < len(lines); i++ {
					if i+1 == len(lines) {
						break // dont go off the end
					}
					split := strings.Fields(lines[i+1])
					if split[0] == "$" {
						break
					}
					if strings.HasPrefix(lines[i+1], "dir") {
						//do nothing for the time being - its an empty dir
					} else {
						file := NewFile(lines[i+1])
						currentDir.files[file.name] = file
					}
				}
			}

		}
	}
	return *root
}
