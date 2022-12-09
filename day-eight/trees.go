package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"
)

//go:embed input-ex.txt
var inputex string

//go:embed input.txt
var input string

func parseGrid(in string) [][]int {
	grid := [][]int{}
	scan := bufio.NewScanner(strings.NewReader(in))
	for i := 0; scan.Scan(); i++ {
		line := scan.Text()
		row := make([]int, len(line), len(line))
		for i, r := range line {
			row[i] = int(r) - 48
		}
		grid = append(grid, row)
	}
	return grid
}

func countEdgeTrees(g [][]int) int {
	return len(g)*2 + len(g[0])*2 - 4
}

func countVisibleTrees(g [][]int) int {
	vis := countEdgeTrees(g)

	for y := 1; y < len(g)-1; y++ {
	cols:
		for x := 1; x < len(g[y])-1; x++ {
			//left
			for l := x - 1; l >= 0; l-- {
				if g[y][l] >= g[y][x] {
					break
				}
				if l == 0 {
					vis++
					continue cols
				}
			}

			//right
			for r := x + 1; r < len(g[y]); r++ {
				if g[y][r] >= g[y][x] {
					break
				}
				if r == len(g[y])-1 {
					vis++
					continue cols
				}
			}

			//up
			for u := y - 1; u >= 0; u-- {
				if g[u][x] >= g[y][x] {
					break
				}
				if u == 0 {
					vis++
					continue cols
				}
			}

			//down
			for d := y + 1; d < len(g); d++ {
				if g[d][x] >= g[y][x] {
					break
				}
				if d == len(g)-1 {
					vis++
				}
			}

		}
	}

	return vis
}

func mapVisScore(g [][]int) map[string]int {
	m := map[string]int{}

	for y := 1; y < len(g)-1; y++ {
		for x := 1; x < len(g[y])-1; x++ {
			key := fmt.Sprintf("%d,%d", x, y)
			var left, right, up, down int

			//left
			for l := x - 1; l >= 0; l-- {
				if g[y][l] >= g[y][x] {
					left++
					break
				}
				left++
			}

			//right
			for r := x + 1; r < len(g[y]); r++ {
				if g[y][r] >= g[y][x] {
					right++
					break
				}
				right++
			}

			//up
			for u := y - 1; u >= 0; u-- {
				if g[u][x] >= g[y][x] {
					up++
					break
				}
				up++
			}

			//down
			for d := y + 1; d < len(g); d++ {
				if g[d][x] >= g[y][x] {
					down++
					break
				}
				down++
			}
			m[key] = left * right * up * down
		}
	}

	return m

}

func main() {

	g := parseGrid(inputex)
	fmt.Println("Part One Ex:", countVisibleTrees(g))
	bestVis := 0
	for _, v := range mapVisScore(g) {
		if v > bestVis {
			bestVis = v
		}
	}
	fmt.Println("Part Two Ex:", bestVis)

	g = parseGrid(input)
	fmt.Println("Part One:", countVisibleTrees(g))
	bestVis = 0
	for _, v := range mapVisScore(g) {

		if v > bestVis {
			bestVis = v
		}
	}
	fmt.Println("Part Two:", bestVis)
}
