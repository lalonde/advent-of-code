package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	_ "embed"
)

const (
	startLabel = "s"
	headLabel  = "H"
	tailLabel  = "T"
	moveRight  = "R"
	moveLeft   = "L"
	moveUp     = "U"
	moveDown   = "D"
)

//go:embed input-ex.txt
var inputex string

//go:embed input.txt
var input string

//go:embed input-exlarge.txt
var inputxl string

type knot struct {
	l *location
}

func (k *knot) distance(o *knot) int {
	x := o.l.x - k.l.x
	y := o.l.y - k.l.y
	return int(math.Sqrt(float64(x*x + y*y)))
}

type rope struct {
	knots []*knot
}

func (r *rope) head() *knot {
	return r.knots[0]
}

type field struct {
	locations map[string]*location
	start     *location
}

func newField() *field {
	f := field{}
	l := newLocation(0, 0, startLabel)
	f.start = l
	f.locations = map[string]*location{l.id: l}
	return &f
}

func (f *field) locationAt(x, y int) *location {
	return f.locations[locationID(x, y)]
}

func (f *field) mustGetLocationAt(x, y int) *location {
	l := f.locationAt(x, y)
	if l == nil {
		return f.addLocation(x, y)
	}
	return l
}

func (f *field) empty(x, y int) bool {
	_, ok := f.locations[locationID(x, y)]
	return !ok
}

func (f *field) addLocation(x, y int, labels ...string) *location {
	if !f.empty(x, y) {
		return nil
	}
	l := newLocation(x, y, labels...)
	f.locations[l.id] = l
	return l
}

type location struct {
	id     string
	x, y   int
	labels map[string]struct{}
}

func newLocation(x, y int, labels ...string) *location {
	labelm := map[string]struct{}{}
	for i := range labels {
		labelm[labels[i]] = struct{}{}
	}

	loc := location{
		id:     locationID(x, y),
		x:      x,
		y:      y,
		labels: labelm,
	}

	return &loc
}

func (l *location) addLabel(x string) {
	l.labels[x] = struct{}{}
}

type ropefield struct {
	r *rope
	f *field
}

func newRopeField(knots int) *ropefield {
	f := newField()
	f.start.addLabel(headLabel)
	f.start.addLabel(tailLabel)

	k := make([]*knot, knots, knots)
	for i := range k {
		k[i] = &knot{f.start}
	}

	return &ropefield{
		r: &rope{k},
		f: f,
	}
}

func (rf *ropefield) moveHeadLeft() {
	rf.moveHeadBy(-1, 0)

}

func (rf *ropefield) moveHeadRight() {
	rf.moveHeadBy(1, 0)
}

func (rf *ropefield) moveHeadUp() {
	rf.moveHeadBy(0, 1)
}

func (rf *ropefield) moveHeadDown() {
	rf.moveHeadBy(0, -1)
}

// TODO: assumption is one move. Should protect
func (rf *ropefield) moveHeadBy(x, y int) {
	from := rf.r.head().l
	to := rf.f.mustGetLocationAt(from.x+x, from.y+y)
	to.addLabel(headLabel)
	rf.r.knots[0].l = to

	for i := 1; i < len(rf.r.knots); i++ {
		leadk := rf.r.knots[i-1]
		followk := rf.r.knots[i]

		yoff := (leadk.l.y - followk.l.y)
		xoff := (leadk.l.x - followk.l.x)
		if followk.distance(leadk) > 1 {
			if yoff != 0 && xoff != 0 {
				followk.l = rf.f.mustGetLocationAt(followk.l.x+posOrNegOne(xoff), followk.l.y+posOrNegOne(yoff))
			} else {
				followk.l = rf.f.mustGetLocationAt(followk.l.x+(xoff/2), followk.l.y+(yoff/2))
			}

			if i == len(rf.r.knots)-1 {
				followk.l.addLabel(tailLabel)
			}
			continue

		}
		break
	}
}

func posOrNegOne(n int) int {
	if n > 0 {
		return 1
	}
	return -1
}

func (rf *ropefield) moveHead(move string, n int) {
	headLoc := rf.r.head().l

	switch move {
	case moveLeft:
		to := headLoc.x - n
		for i := headLoc.x - 1; i >= to; i-- {
			rf.moveHeadLeft()
		}
	case moveRight:
		to := headLoc.x + n
		for i := headLoc.x + 1; i <= to; i++ {
			rf.moveHeadRight()
		}
	case moveUp:
		to := headLoc.y + n
		for i := headLoc.y + 1; i <= to; i++ {
			rf.moveHeadUp()
		}
	case moveDown:
		to := headLoc.y - n
		for i := headLoc.y - 1; i >= to; i-- {
			rf.moveHeadDown()
		}

	}
}

func (rf *ropefield) vis() {
	var minx, maxx, miny, maxy int
	for _, v := range rf.f.locations {
		if v.x < minx {
			minx = v.x
		}
		if v.x > maxx {
			maxx = v.x
		}
		if v.y < miny {
			miny = v.y
		}
		if v.y > maxy {
			maxy = v.y
		}
	}

	rows := maxy - miny + 1
	cols := maxx - minx + 1

	g := make([][]string, rows, rows)
	for i := range g {
		g[i] = make([]string, cols, cols)
		for j := range g[i] {
			g[i][j] = "."
		}
	}

	for i, k := range rf.r.knots {
		l := k.l
		g[l.y-miny][l.x-minx] = strconv.Itoa(i)
	}

	for y := len(g) - 1; y >= 0; y-- {
		for x := range g[y] {
			fmt.Print(g[y][x])
		}
		fmt.Print("\n")
	}

	fmt.Println("")
}

func parseAndRunFor(part, in string, knots int, vis bool) {
	motions := bufio.NewScanner(strings.NewReader(in))
	area := newRopeField(knots)
	for motions.Scan() {
		motion := strings.Fields(motions.Text())
		n, err := strconv.Atoi(motion[1])
		if err != nil {
			panic(err.Error())
		}
		area.moveHead(motion[0], n)
		if vis {
			fmt.Println(motion)
			area.vis()
		}
	}

	tailVisits := 0
	for _, v := range area.f.locations {
		if _, ok := v.labels[tailLabel]; ok {
			tailVisits++
		}
	}

	fmt.Println(part, tailVisits)

}

func locationID(x, y int) string {
	return fmt.Sprintf("x:%d,y:%d", x, y)
}

func main() {
	parseAndRunFor("Part One Ex", inputex, 2, false)
	parseAndRunFor("Part One", input, 2, false)
	parseAndRunFor("Part Two Ex", inputxl, 10, false)
	parseAndRunFor("Part Two", input, 10, false)
}
