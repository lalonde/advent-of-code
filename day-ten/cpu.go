package main_test

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input-ex.txt
var inputex string

//go:embed input.txt
var input string

type clock struct {
	cycle      int
	components []chan<- int
}

func (cl *clock) connect() <-chan int {
	ch := make(chan int)
	cl.components = append(cl.components, ch)
	return ch
}

func (cl *clock) tick() {
	cl.cycle++
	for i := range cl.components {
		cl.components[i] <- cl.cycle
	}
}

type cpu struct {
	clock <-chan int
	x     int
	trace map[int]int
}

func (c *cpu) noop() {
	cy := <-c.clock
	fmt.Println("cpu", cy)
	c.trace[cy] = c.x
}

func (c *cpu) addx(n int) {
	cy := <-c.clock
	fmt.Println("cpu", cy)
	c.trace[cy] = c.x
	cy = <-c.clock
	fmt.Println("cpu", cy)
	c.trace[cy] = c.x
	c.x += n
}

type crt struct {
	clock  <-chan int
	reg    *int
	screen [][]string
}

func (c *crt) on() {
	for {
		cycle := <-c.clock
		fmt.Println("crt", cycle)
		hpos := 40
		if x := cycle % 40; x != 0 {
			hpos = x
		}

		if hpos <= *c.reg && hpos >= *c.reg {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		if hpos == 40 {
			fmt.Println("")
		}
	}
}

type com struct {
	cpu   *cpu
	clock *clock
	crt   *crt
	on    bool
}

func newCom() *com {
	cl := &clock{
		cycle: 0,
	}
	c := &com{
		clock: cl,
		cpu: &cpu{
			clock: cl.connect(),
			x:     1,
			trace: make(map[int]int),
		},
		crt: &crt{
			clock: cl.connect(),
		},
	}
	c.crt.reg = &c.cpu.x
	return c
}

func (c *com) powerOn() {
	if c.on {
		return
	}

	c.on = true

	go func() {
		for c.on {
			c.clock.tick()
		}
	}()
	go c.crt.on()
}

func (c *com) powerOff() {
	c.on = false
}

func (c *com) run(r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.Fields(s.Text())
		switch l[0] {
		case "noop":
			c.cpu.noop()
		case "addx":
			n, err := strconv.Atoi(l[1])
			if err != nil {
				panic(err.Error())
			}
			c.cpu.addx(n)
		}
	}
}

func main() {
	partOne("Part One Ex:", inputex)
	//	partOne("Part One:", input)
}

func partOne(name, in string) {
	c := newCom()
	c.powerOn()
	c.run(strings.NewReader(in))
	c.powerOff()

	s := 0
	x := []int{20, 60, 100, 140, 180, 220}

	for i := range x {
		s += x[i] * c.cpu.trace[x[i]]
	}
	fmt.Println(name, s)
}
