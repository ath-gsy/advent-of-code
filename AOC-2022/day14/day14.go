package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func buildPairs(pair string) (int, int) {
	xy := strings.Split(pair, ",")
	X, _ := strconv.Atoi(xy[0])
	Y, _ := strconv.Atoi(xy[1])
	return X, Y
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Pos struct {
	x int
	y int
}

func main() {
	input := strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	grid := make(map[Pos]byte) //byte == uint8 ; aussi création avec grid := map[Pos]byte{}
	for _, line := range lines {
		pairs := strings.Split(line, "->")
		for i := 0; i < len(pairs)-1; i++ {
			// xy := strings.Split(pair, ",")
			// X, _ := strconv.Atoi(xy[0])
			// Y, _ := strconv.Atoi(xy[1])
			x1, y1 := buildPairs(pairs[i])
			x2, y2 := buildPairs(pairs[i+1])
			for a := min(x1, x2); a <= max(x1, x2); a++ {
				for b := min(y1, y2); b <= max(y1, y2); b++ {
					grid[Pos{a, b}] = '#'

				}
			}
		}
	}
}

//// VOIR CODE GITHUB
// func display(grid map[Pos]byte) {
// 	xmin := math.MaxInt
// 	xmax := math.MinInt
// 	ymin := math.MaxInt
// 	ymax := math.MinInt
// 	for i :=0
// }
