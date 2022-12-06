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

//go:embed input_example.txt
var inputexample string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(inputexample))
	stacks := parseCrateStack(scanner)
	moveCrates(scanner, stacks)

	fmt.Println("Example:", stackHeads(stacks))

	scanner = bufio.NewScanner(strings.NewReader(input))
	stacks = parseCrateStack(scanner)
	moveCrates(scanner, stacks)

	fmt.Println("Part One:", stackHeads(stacks))

	scanner = bufio.NewScanner(strings.NewReader(inputexample))
	stacks = parseCrateStack(scanner)
	moveCrates9001(scanner, stacks)

	fmt.Println("Example Two:", stackHeads(stacks))

	scanner = bufio.NewScanner(strings.NewReader(input))
	stacks = parseCrateStack(scanner)
	moveCrates9001(scanner, stacks)

	fmt.Println("Part Two:", stackHeads(stacks))

}

func moveCrates(scanner *bufio.Scanner, stacks []*stack[rune]) {
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		n := mustAtoi(moves[1])
		from := stacks[mustAtoi(moves[3])-1]
		to := stacks[mustAtoi(moves[5])-1]

		for i := 1; i <= n; i++ {
			to.push(from.pop())
		}
	}
}

func moveCrates9001(scanner *bufio.Scanner, stacks []*stack[rune]) {
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		n := mustAtoi(moves[1])
		from := stacks[mustAtoi(moves[3])-1]
		to := stacks[mustAtoi(moves[5])-1]

		buf := from.popn(n)

		for i := n; i != 0; i-- {
			to.push(buf[i-1])
		}
	}
}

func stackHeads(stacks []*stack[rune]) string {
	r := ""
	for i := range stacks {
		x := "_"
		if !stacks[i].isEmpty() {
			x = string(stacks[i].peek())
		}
		r = r + x
	}
	return r
}

func parseCrateStack(scanner *bufio.Scanner) []*stack[rune] {
	buf := make([]string, 0)

	for scanner.Scan() && scanner.Text() != "" {
		if scanner.Text() == "" {
			break
		}
		buf = append(buf, scanner.Text())
	}

	index := map[int]int{}
	indexLine := buf[len(buf)-1]
	stacks := make([]*stack[rune], 0)
	for i := 1; true; i++ {
		x := strings.Index(indexLine, strconv.Itoa(i))
		if x == -1 {
			break
		}
		index[i] = x
		stacks = append(stacks, &stack[rune]{})
	}

	columns := len(stacks)

	for i := len(buf) - 2; i >= 0; i-- {
		for j := 1; j <= columns; j++ {
			z := buf[i][index[j]]
			if z >= 65 && z <= 90 {
				stacks[j-1].push(rune(z))
			}
		}
	}
	return stacks
}

func mustAtoi(in string) int {
	x, err := strconv.Atoi(in)
	if err != nil {
		panic(err.Error())
	}
	return x
}
