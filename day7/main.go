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
	// not  1405252 - too low
	// wft  1478292
	//Al's  1749646
	//        71943

}

func part1() int {
	root := structure(fruitpicker.ReadInput())
	if debug {
		//root.print(0)
	}
	sizes := make(map[string]int)
	root.calculateSizes(sizes)
	total := 0
	for k := range sizes {
		small := ""
		if sizes[k] <= 100000 {
			total += sizes[k]
			small = " <==="
		}
		if debug {
			fmt.Printf("%v %v%v\n", k, sizes[k], small)
		}
	}
	s := sizes["/"]
	fmt.Println(s) // total size
	//  48381165
	fmt.Println(total)
	return total
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

func (d *dir) calculateSizes(m map[string]int) int {
	total := 0
	for _, v := range d.dirs {
		total += v.calculateSizes(m)
	}

	for _, f := range d.files {
		total += f.size
	}
	_, ok := m[d.name]
	if ok {
		panic("duplicate filename!!")
	}
	m[d.name] = total
	return total
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
		fmt.Printf(" %v ", i)
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
					fmt.Printf(" %v ", i+1)
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
	fmt.Printf("\n")

	return *root
}
