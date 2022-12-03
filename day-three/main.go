package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type rucksack string

func (r rucksack) contents() string {
	return string(r)
}

func (r rucksack) compartmentDups() []rune {
	dups := []rune{}
	a, b := r.compartmentContents()

	amap := map[rune]bool{}
	for _, r := range a {
		amap[r] = false
	}

	for _, r := range b {
		if x, ok := amap[r]; ok && !x {
			amap[r] = true
			dups = append(dups, r)
		}
	}

	fmt.Println("Dups in", string(r), dups)
	return dups
}

func (r rucksack) compartmentContents() (cOne, cTwo string) {
	if len(r.contents())%2 == 1 {
		panic("Uneven rucksack contents! The game is rigged!")
	}
	c := r.contents()
	p := len(c) / 2
	cOne = c[:p]
	cTwo = c[p:]
	return
}

func (r rucksack) itemCounts() map[rune]int {
	counts := map[rune]int{}
	for _, v := range r.contents() {
		counts[v]++
	}
	return counts
}

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var total int
	badges := []rune{}
	groupItemCount := map[rune]int{}

	for i := 1; scanner.Scan(); i++ {
		sack := rucksack(scanner.Text())
		for _, item := range sack.compartmentDups() {
			total += itemPriority(item)
		}

		//TODO: scan contents once
		for k := range sack.itemCounts() {
			groupItemCount[k]++
		}

		if i%3 == 0 {
			fmt.Println(groupItemCount)
			group := i / 3
			for k, v := range groupItemCount {
				if v == 3 {
					fmt.Println("Group", group, "Badge", string(k))
					badges = append(badges, k)
				}
			}
			groupItemCount = map[rune]int{}
		}
	}

	fmt.Println("Part One:", total)
	fmt.Println("Part Two:", sumItemPriority(badges))
}

func sumItemPriority(items []rune) int {
	var sum int
	for _, r := range items {
		sum += itemPriority(r)
	}
	return sum
}

func itemPriority(item rune) int {
	if item >= 'a' {
		return int(item - 96)
	}
	return int(item - 38)
}
