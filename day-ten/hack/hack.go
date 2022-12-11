package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input-ex.txt
var inputex string

//go:embed input.txt
var input string

func main() {

	scanner := bufio.NewScanner(strings.NewReader(input))

	cycle := 0
	regx := 1

	for scanner.Scan() {
		cycle++
		crt(cycle, regx)
		line := strings.Fields(scanner.Text())
		if line[0] == "addx" {
			cycle++
			regx += mustAtoi(line[1])
			crt(cycle, regx)

		}
	}
}

func crt(cycle, reg int) {
	hpos := 40
	if x := cycle % 40; x != 0 {
		hpos = x
	}

	if hpos <= reg+1 && hpos >= reg-1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if hpos == 40 {
		fmt.Println("")
	}
}

func mustAtoi(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return x
}
