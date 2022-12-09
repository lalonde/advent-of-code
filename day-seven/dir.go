package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

//go:embed input-example.txt
var inputex string

type fsentry struct {
	name   string
	parent *fsentry
	items  map[string]*fsentry
	size   int
	isdir  bool
}

func (f *fsentry) fullpath() string {
	path := f.name
	for p := f.parent; p.parent != nil; p = p.parent {
		path = p.name + "/" + path
	}
	return path
}

type term struct {
	root     *fsentry
	currdir  *fsentry
	dooutput bool
}

func newterm() *term {
	root := &fsentry{name: "/"}
	root.items = map[string]*fsentry{}
	return &term{root, root, false}
}

func (t *term) pwd() string {
	return t.currdir.fullpath()
}

func (t *term) mkdir(name string) *fsentry {
	return t.newOrVistFE(name, 0, true)
}

func (t *term) touch(name string, size int) *fsentry {
	return t.newOrVistFE(name, size, false)
}

func (t *term) newOrVistFE(name string, size int, isdir bool) *fsentry {
	if fe, ok := t.currdir.items[name]; ok {
		return fe
	}

	fe := &fsentry{
		name:   name,
		parent: t.currdir,
		isdir:  isdir,
		size:   size,
	}

	if isdir {
		fe.items = map[string]*fsentry{}
	}

	t.currdir.items[name] = fe

	return fe
}

func (t *term) processinput(in *bufio.Scanner) {
	for in.Scan() {
		l := strings.Fields(in.Text())
		if len(l) == 0 {
			continue
		}

		if l[0] == "$" {
			t.dooutput = false
			t.doCmd(l[1:])
			continue
		}

		if t.dooutput {
			t.handleOutput(l)
		}
	}
}

func (t *term) cd(name string) {
	if name == ".." {
		t.currdir = t.currdir.parent
		return
	}
	t.currdir = t.mkdir(name)
}

func (t *term) doCmd(cmd []string) {
	switch cmd[0] {
	case "ls":
		t.dooutput = true
	case "cd":
		t.cd(cmd[1])
	}
}

func (t *term) handleOutput(out []string) {
	if out[0] == "dir" {
		t.mkdir(out[1])
		return
	}
	size, err := strconv.Atoi(out[0])
	if err != nil {
		fmt.Println("cant handle", out)
	}
	t.touch(out[1], size)
}

func main() {
	totalspace := 70000000
	spaceneeded := 30000000

	s := bufio.NewScanner(strings.NewReader(inputex))
	t := newterm()
	t.processinput(s)
	fmt.Println("Example: ", t.sumDirsUnder100k())

	for k, v := range t.mapDirSize() {
		fmt.Println(k, v)
	}

	ds := t.mapDirSize()
	fsfree := totalspace - ds["/"]

	fmt.Println("Free space:", fsfree)

	minsize := totalspace
	for _, v := range ds {
		ifdel := fsfree + v
		if ifdel >= spaceneeded && ifdel < fsfree+minsize {
			minsize = v
		}
	}

	fmt.Println("min dir size:", minsize)

	s = bufio.NewScanner(strings.NewReader(input))
	t = newterm()
	t.processinput(s)
	fmt.Println("Part One: ", t.sumDirsUnder100k())

	ds = t.mapDirSize()
	minsize = totalspace
	fsfree = totalspace - ds["/"]
	for k, v := range ds {
		ifdel := fsfree + v
		if ifdel >= spaceneeded && ifdel < fsfree+minsize {
			fmt.Println(k, v)
			minsize = v
		}
	}

	fmt.Println("Part Two:", minsize)

}

func (entry *fsentry) walk(f func(*fsentry)) {
	for _, e := range entry.items {
		f(e)
		e.walk(f)
	}
}

func (t *term) mapDirSize() map[string]int {
	acc := map[string]int{}

	t.root.walk(func(item *fsentry) {
		for p := item.parent; p.parent != nil; p = p.parent {
			acc[p.fullpath()] += item.size
		}
	})

	return acc
}

func (t *term) sumDirsUnder100k() int {
	m := t.mapDirSize()
	x := 0
	for _, v := range m {
		if v <= 100000 {
			x += v
		}

	}
	return x
}
