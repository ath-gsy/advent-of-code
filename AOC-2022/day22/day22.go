package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Position struct {
	x         int
	y         int
	direction int
}

// /Position methods
func (position *Position) turn(RorL string) { //quasiment fait que par github copilot o.O
	// fmt.Println(position.direction)
	if RorL == "R" {
		position.direction = (position.direction + 1) % 4
	} else if RorL == "L" {
		position.direction = (position.direction + 3) % 4
	}
}
func (position *Position) move1() { //quasiment fait que par github copilot o.O

	if !position.void() {
		switch position.direction {
		case 0:
			position.y++
		case 1:
			position.x++
		case 2:
			position.y--
		case 3:
			position.x--
		}
	} else {
		nextx, nexty := position.getNextPoint()
		position.x = nextx
		position.y = nexty

	}
}
func (position *Position) detectObstacle() bool {

	nextx, nexty := position.getNextPoint()

	if mapPlan[nextx][nexty] == "#" {
		return true
	} else {
		return false
	}
	// switch position.direction {
	// case 0:
	// 	if mapPlan[position.x][position.y+1] == "#" {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// case 1:
	// 	if mapPlan[position.x+1][position.y] == "#" {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// case 2:
	// 	if mapPlan[position.x][position.y-1] == "#" {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// case 3:
	// 	if mapPlan[position.x-1][position.y] == "#" {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// }
}
func (position *Position) void() bool { //quasiment TOUT fait par github copilot wow o.O
	switch position.direction {
	case 0:
		if mapPlan[position.x][position.y+1] == " " {
			return true
		} else {
			return false
		}
	case 1:
		if mapPlan[position.x+1][position.y] == " " {
			return true
		} else {
			return false
		}
	case 2:
		if mapPlan[position.x][position.y-1] == " " {
			return true
		} else {
			return false
		}
	case 3:
		if mapPlan[position.x-1][position.y] == " " {
			return true
		} else {
			return false
		}
	}
	return false
}
func (position *Position) getNextPoint() (int, int) {
	if position.void() {
		switch position.direction {
		case 0:
			return connectedPointsHorizontal[position.x].x1, connectedPointsHorizontal[position.x].y1
		case 2:
			return connectedPointsHorizontal[position.x].x2, connectedPointsHorizontal[position.x].y2
		case 1:
			return connectedPointsVertical[position.y].x1, connectedPointsVertical[position.y].y1
		case 3:
			return connectedPointsVertical[position.y].x2, connectedPointsVertical[position.y].y2
		}
	} else {
		switch position.direction {
		case 0:
			return position.x, position.y + 1
		case 1:
			return position.x + 1, position.y
		case 2:
			return position.x, position.y - 1
		case 3:
			return position.x - 1, position.y
		}
	}
	return position.x, position.y

}

type PointsPair struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func printMap() {
	f, err := os.Create("mapPlan.ext") // création d'un fichier .ext car avast m'embête avec les fichiers .txt
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	for _, line := range mapPlan {
		f.WriteString(strings.Join(line, "") + "\n")
	}
}

var (
	//go:embed input.txt
	input string

	instructions []string
	mapPlan      [][]string

	currentPosition Position

	connectedPointsHorizontal = make(map[int]PointsPair)
	connectedPointsVertical   = make(map[int]PointsPair)
)

func init() {
	input := strings.TrimSuffix(input, "\n")
	content := strings.Split(input, "\n\n")

	//create mapPlan
	mapContent := strings.Split(content[0], "\n")
	i := 0
	line0 := make([]string, len(mapContent[0])+2)
	for i < len(mapContent[0])+2 {
		line0[i] = " "
		i++
	}
	mapPlan = append(mapPlan, line0)

	tmpLine := make([]string, 0)
	for _, line := range mapContent {
		tmpLine = append(tmpLine, " ")
		tmpLine = append(tmpLine, strings.Split(line, "")...)
		if len(tmpLine) < len(mapContent[0])+2 {
			for i := len(tmpLine); i < len(mapContent[0])+1; i++ {
				tmpLine = append(tmpLine, " ")
			}
		}
		tmpLine = append(tmpLine, " ")
		mapPlan = append(mapPlan, tmpLine)
		tmpLine = nil
	}
	i = 0
	line0 = make([]string, len(mapContent[0])+2)
	for i < len(mapContent[0])+2 {
		line0[i] = " "
		i++
	}
	mapPlan = append(mapPlan, line0)
	////////////////////////

	//create instructions
	rawInstructions := strings.Split(content[1], "\n")
	rawInstructions = strings.Split(rawInstructions[0], "")

	tmpBuffer := ""
	for _, instructionItem := range rawInstructions {
		_, err := strconv.Atoi(instructionItem)
		if err == nil {
			tmpBuffer += instructionItem
		} else {
			instructions = append(instructions, tmpBuffer)
			tmpBuffer = ""
			instructions = append(instructions, instructionItem)
		}
	}
	instructions = append(instructions, tmpBuffer)
	////////////////////////

	//find start position
	ystart := 0
	for i, point := range mapPlan[1] {
		if point == "." {
			ystart = i
			break
		}
	}
	currentPosition = Position{1, ystart, 0}
	////////////////////////

	//create connectedPoints horizontal and vertical
	for i, line := range mapPlan {
		y1 := 0
		y2 := 0
		for j, point := range line {
			if point != " " {
				y1 = j
				break
			}
		}
		for j := y1; j < len(line); j++ {
			if line[j] == " " {
				y2 = j - 1
				break
			}
		}
		if y1 != y2+1 {
			// mapPlan[i][y1] = "-"
			// mapPlan[i][y2] = "-"
			connectedPointsHorizontal[i] = PointsPair{i, y1, i, y2}
		}
	}
	for j := range mapPlan[0] {
		x1 := 0
		x2 := 0
		for i, _ := range mapPlan {
			if mapPlan[i][j] != " " {
				x1 = i
				break
			}
		}
		for i := x1; i < len(mapPlan); i++ {
			if mapPlan[i][j] == " " {
				x2 = i - 1
				break
			}
		}
		if x1 != x2+1 {
			// mapPlan[x1][j] = "|"
			// mapPlan[x2][j] = "|"
			connectedPointsVertical[j] = PointsPair{x1, j, x2, j}
		}
	}
	////////////////////////
}

func updateMap() {
	switch currentPosition.direction {
	case 0:
		mapPlan[currentPosition.x][currentPosition.y] = ">"
	case 1:
		mapPlan[currentPosition.x][currentPosition.y] = "v"
	case 2:
		mapPlan[currentPosition.x][currentPosition.y] = "<"
	case 3:
		mapPlan[currentPosition.x][currentPosition.y] = "^"
	}

}

func part1() {
	updateMap()
	// printMap()

	for id, instruction := range instructions {
		if id%5 == 0 {
			// fmt.Println("--------------------------------")
		}
		// fmt.Print(instruction, " || ")
		if instruction == "R" || instruction == "L" {
			currentPosition.turn(instruction)
			updateMap()
		} else {
			distance, _ := strconv.Atoi(instruction)
			i := 0
			for i < distance {
				if currentPosition.detectObstacle() {
					break
				}
				currentPosition.move1()
				updateMap()
				i++
			}

		}
		// fmt.Println("current state: ", currentPosition.x, currentPosition.y, currentPosition.direction, "; ")

		// printMap()
	}

	result := 1000*(currentPosition.x) + 4*(currentPosition.y) + currentPosition.direction
	fmt.Println("result: ", result)
}

func main() {

	start := time.Now()
	fmt.Print("\npart1: ")
	part1()
	fmt.Println(time.Since(start))
	printMap()

}
