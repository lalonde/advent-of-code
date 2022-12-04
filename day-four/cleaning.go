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

type assignment struct {
	a, b int
}

func parseAssignmentPair(in string) (assignment, assignment) {
	x := strings.Split(in, ",")

	a, b := parseAssignment(x[0])
	a1 := assignment{
		a, b,
	}
	a, b = parseAssignment(x[1])
	a2 := assignment{
		a, b,
	}
	return a1, a2
}

func parseAssignment(in string) (int, int) {
	x := strings.Split(in, "-")
	a, err := strconv.Atoi(x[0])
	if err != nil {
		panic("Bad input on assignment")
	}
	b, err := strconv.Atoi(x[1])
	if err != nil {
		panic("Bad input on assignment")
	}
	return a, b
}

func (a assignment) contains(other assignment) bool {
	if a.a <= other.a && a.b >= other.b {
		return true
	}
	return false
}

func (a assignment) overlaps(other assignment) bool {
	if a.a >= other.a && a.a <= other.b {
		return true
	}

	if a.b >= other.a && a.b <= other.b {
		return true
	}

	return false
}

func main() {
	var fullOverlap int
	var overlap int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		a1, a2 := parseAssignmentPair(scanner.Text())

		if a1.contains(a2) || a2.contains(a1) {
			fullOverlap++
		}
		//TODO: Really thought one overlap call would work
		// comeback to see why two calls were needed.
		if a1.overlaps(a2) || a2.overlaps(a1) {
			overlap++
		}
	}

	fmt.Println("Part One:", fullOverlap)
	fmt.Println("Part Two:", overlap)
}
