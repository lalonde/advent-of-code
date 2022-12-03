package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"
)

const (
	rock = iota + 1
	paper
	scisors

	win  = 6
	loss = 0
	draw = 3
)

//go:embed 2-1-input.txt
var input string

const testinput = `A Y
B X
C Z`

func main() {
	fmt.Println(playGame(input))
}

func playGame(input string) int {
	score := 0

	scan := bufio.NewScanner(strings.NewReader(input))

	for scan.Scan() {
		l := scan.Text()
		if len(l) != 3 {
			continue
		}

		theirs := decodeChoice(rune(l[0]))
		result := decodeResult(rune(l[2]))

		// Part one
		//		mine := decodeChoice(rune(l[2]))

		//		score += scoreRound(theirs, mine)
		score += choiceForOutcome(theirs, result) + result
	}

	return score
}

func choiceForOutcome(theirs, result int) int {
	if result == draw {
		return theirs
	}

	if result == win {
		winner := theirs + 1
		if winner > 3 {
			winner -= 3
		}
		return winner
	}

	if result == loss {
		loser := theirs + 2
		if loser > 3 {
			loser -= 3
		}
		return loser
	}

	panic("no match")
}

func scoreRound(theirs, mine int) int {
	if theirs == mine {
		return mine + draw
	}

	if mine == rock && theirs == scisors {
		return mine + win
	}

	if mine == paper && theirs == rock {
		return mine + win
	}

	if mine == scisors && theirs == paper {
		return mine + win
	}

	return mine
}

func decodeChoice(input rune) int {
	switch input {
	case 'A', 'X':
		return rock
	case 'B', 'Y':
		return paper
	case 'C', 'Z':
		return scisors
	default:
		return 0
	}
}

func decodeResult(input rune) int {
	switch input {
	case 'X':
		return loss
	case 'Y':
		return draw
	case 'Z':
		return win
	}
	panic("no match")
}
