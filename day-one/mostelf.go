package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed 1-1-input.txt
var input string

func main() {
	fmt.Println(findMostElf(input))
}

func findMostElf(in string) (int, int, int) {
	var mostElf, mostElfCal, currentElf, currentElfCal int
	currentElf = 1

	var top3 []int = make([]int, 3, 3)

	scanner := bufio.NewScanner(strings.NewReader(in))

	for scanner.Scan() {
		l := scanner.Text()

		if strings.TrimSpace(l) == "" {
			if currentElfCal >= mostElfCal {
				mostElf = currentElf
				mostElfCal = currentElfCal
			}
			doTop3(currentElfCal, top3)
			currentElf++
			currentElfCal = 0

			fmt.Println(top3)
			continue
		}

		if c, e := strconv.Atoi(l); e == nil {
			currentElfCal += c
		}
	}

	if currentElfCal >= mostElfCal {
		mostElf = currentElf
		mostElfCal = currentElfCal
		doTop3(currentElfCal, top3)
	}

	sum3 := 0
	for i := range top3 {
		sum3 += top3[i]
	}

	return mostElf, mostElfCal, sum3
}

func doTop3(n int, s []int) {
	for x := 0; x < len(s); x++ {
		if n > s[x] {
			for y := len(s) - 1; y > x; y-- {
				s[y] = s[y-1]
			}
			s[x] = n
			break
		}
	}
}
