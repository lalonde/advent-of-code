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
	id string
	l  *location
}

type rope struct {
	head, tail *knot
}

func (r *rope) headLoc() *location {
	return r.head.l
}

func (r *rope) tailLoc() *location {
	return r.tail.l
}

func (r *rope) headTailDistance() int {
	x := r.head.l.x - r.tail.l.x
	y := r.head.l.y - r.tail.l.y
	return int(math.Sqrt(float64(x*x + y*y)))
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

func (f *field) empty(x, y int) bool {
	_, ok := f.locations[locationID(x, y)]
	return !ok
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

func (f *field) addLocation(x, y int, labels ...string) *location {
	if !f.empty(x, y) {
		return nil
	}
	l := newLocation(x, y, labels...)
	f.locations[l.id] = l
	return l
}

func (f *field) lookWestOf(x, y int) *location {
	return f.locationAt(x-1, y)
}

func (f *field) lookEastOf(x, y int) *location {
	return f.locationAt(x+1, y)
}

func (f *field) lookNorthOf(x, y int) *location {
	return f.locationAt(x, y+1)
}

func (f *field) lookSouthOf(x, y int) *location {
	return f.locationAt(x, y-1)
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

func newRopeField() *ropefield {
	f := newField()
	f.start.addLabel(headLabel)
	f.start.addLabel(tailLabel)
	r := rope{
		head: &knot{headLabel, f.start},
		tail: &knot{tailLabel, f.start},
	}
	return &ropefield{
		r: &r,
		f: f,
	}
}

func (rf *ropefield) moveHeadLeft() {
	l := rf.r.head.l
	rf.moveHeadTo(l.x-1, l.y)

}

func (rf *ropefield) moveHeadRight() {
	l := rf.r.head.l
	rf.moveHeadTo(l.x+1, l.y)
}

func (rf *ropefield) moveHeadUp() {
	l := rf.r.head.l
	rf.moveHeadTo(l.x, l.y+1)
}

func (rf *ropefield) moveHeadDown() {
	l := rf.r.head.l
	rf.moveHeadTo(l.x, l.y-1)
}

// TODO: assumption is one move. Should protect
func (rf *ropefield) moveHeadTo(x, y int) {
	from := rf.r.head.l
	to := rf.f.mustGetLocationAt(x, y)
	to.addLabel(headLabel)
	rf.r.head.l = to
	if rf.r.headTailDistance() > 1 {
		rf.r.tail.l = from
		from.addLabel(tailLabel)
	}
}

func (rf *ropefield) moveHead(move string, n int) {
	switch move {
	case moveLeft:
		to := rf.r.headLoc().x - n
		for i := rf.r.headLoc().x - 1; i >= to; i-- {
			rf.moveHeadLeft()
		}
	case moveRight:
		to := rf.r.headLoc().x + n
		for i := rf.r.headLoc().x + 1; i <= to; i++ {
			rf.moveHeadRight()
		}
	case moveUp:
		to := rf.r.headLoc().y + n
		for i := rf.r.headLoc().y + 1; i <= to; i++ {
			rf.moveHeadUp()
		}
	case moveDown:
		to := rf.r.headLoc().y - n
		for i := rf.r.headLoc().y - 1; i >= to; i-- {
			rf.moveHeadDown()
		}

	}
}

func main() {
	parseAndRunFor("Part One Ex", inputex)
	parseAndRunFor("Part One", input)

}

func parseAndRunFor(part, in string) {
	motions := bufio.NewScanner(strings.NewReader(in))
	area := newRopeField()
	for motions.Scan() {
		motion := strings.Fields(motions.Text())
		n, err := strconv.Atoi(motion[1])
		if err != nil {
			panic(err.Error())
		}
		area.moveHead(motion[0], n)
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
