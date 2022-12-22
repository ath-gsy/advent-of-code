package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

type Position struct {
	x         int
	y         int
	direction int
}

func (position Position) turn(RorL string) {
	if RorL == "R" {
		position.direction = (position.direction + 1) % 4
	} else {
		position.direction = (position.direction - 1) % 4
	}
}

type PointsPair struct {
	connected bool
	x1        int
	y1        int
	x2        int
	y2        int
}

type Instruction struct {
	number        int
	monkey1       string
	monkey1Number int
	operation     string
	monkey2       string
	monkey2Number int
}

var (
	//go:embed input.txt
	input string

	mapContours []int

	currentPosition Position

	connectedPointsHorizontal = make(map[int]PointsPair)
	connectedPointsVertical   = make(map[int]PointsPair)
)

func init() {
	input := strings.TrimSuffix(input, "\n")
	content := strings.Split(input, "\n\n")

	mapContent := strings.Split(content[0], "\n")
	instructions := strings.Split(content[1], "\n")

	for _, line := range mapContent {

	}

}

func part1() {

}

func main() {

	start := time.Now()
	fmt.Print("part1: ")
	part1()
	fmt.Println(time.Since(start))

}
