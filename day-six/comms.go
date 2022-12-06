package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input-example.txt
var examples string

//go:embed input-example2.txt
var examples2 string

//go:embed input.txt
var input string

func main() {
	scnr := bufio.NewScanner(strings.NewReader(examples))

	fmt.Println("Examples:")
	for scnr.Scan() {
		fmt.Println(scnr.Text(), "starts at", findCommStart(scnr.Text()))
	}

	fmt.Println("Part One:", findCommStart(input))

	scnr = bufio.NewScanner(strings.NewReader(examples2))
	fmt.Println("Example2:")
	for scnr.Scan() {
		fmt.Println(scnr.Text(), "starts at", findMsgStart(scnr.Text()))
	}

	// This seems odd. What if the msg start is before the comm start...
	// The spec was a bit vauge.
	// Either way, this produced the correct answer.
	fmt.Println("Part Two:", findMsgStart(input))
}

// Looks for a sequence of 4 chars where each is unique
// Returns -1 if not found
func findCommStart(in string) int {
	return posAfterNUnique(in, 4)
}

func findMsgStart(in string) int {
	return posAfterNUnique(in, 14)
}

// Looks for a sequence of n unique chars.
// Returns the position after sequence is identified
func posAfterNUnique(in string, n int) int {
	if len(in) < n || n < 1 {
		return -1
	}

	for i := n - 1; i < len(in); i++ {
		m := map[rune]int{}
		for j := 0; j < n; j++ {
			m[rune(in[i-j])]++
		}

		if len(m) == n {
			return i + 1
		}
	}

	return -1

}
