package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func minArray(array []int) int {
	min := array[0]
	for _, value := range array {
		if value < min {
			min = value
		}
	}
	return min
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mapPlanPrinter(mapPlan [][]string, theExpeditionPosition Expedition) {
	f, err := os.Create("mapPlan.ext") // création d'un fichier .ext car avast m'embête avec les fichiers .txt
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	mapPlan[theExpeditionPosition.position.x][theExpeditionPosition.position.y] = "E"

	for _, line := range mapPlan {
		f.WriteString(strings.Join(line, "") + "\n")
	}
	mapPlan[theExpeditionPosition.position.x][theExpeditionPosition.position.y] = "."
}
func mapWindPrinter(mapPlan [][]string) {
	f, err := os.Create("mapPlan.ext") // création d'un fichier .ext car avast m'embête avec les fichiers .txt
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	for _, line := range mapPlan {
		f.WriteString(strings.Join(line, "") + "\n")
	}
}

func mapPlanCreator(newWindPositions map[Position][]Wind, newMapPlan [][]string) [][]string {
	for i := range newMapPlan {
		newMapPlan[i] = make([]string, len(mapPlanModel[i]))
		copy(newMapPlan[i], mapPlanModel[i])
	}
	for _, winds := range newWindPositions {
		if len(winds) > 1 {
			newMapPlan[winds[0].position.x][winds[0].position.y] = strconv.Itoa(len(winds))
		} else {
			switch winds[0].direction {
			case 0:
				newMapPlan[winds[0].position.x][winds[0].position.y] = ">"
			case 1:
				newMapPlan[winds[0].position.x][winds[0].position.y] = "v"
			case 2:
				newMapPlan[winds[0].position.x][winds[0].position.y] = "<"
			case 3:
				newMapPlan[winds[0].position.x][winds[0].position.y] = "^"
			}
		}
	}

	return newMapPlan
}

type Position struct {
	x int
	y int
}

type Wind struct {
	position  Position
	direction int
}

func (wind *Wind) detectObstacle(mapPlan [][]string) bool {
	switch wind.direction {
	case 0:
		if mapPlan[wind.position.x][wind.position.y+1] == "#" {
			return true
		} else {
			return false
		}
	case 1:
		if mapPlan[wind.position.x+1][wind.position.y] == "#" || wind.position.x+1 == len(mapPlan)-1 {
			//après vérification, en fait l'input est bien construite : cad sans vent ascendant ou descendant sur la première ou dernière colonne
			return true
		} else {
			return false
		}
	case 2:
		if mapPlan[wind.position.x][wind.position.y-1] == "#" {
			return true
		} else {
			return false
		}
	case 3:
		if mapPlan[wind.position.x-1][wind.position.y] == "#" || wind.position.x-1 == 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

type TimeFrame struct {
	windPositions      map[Position][]Wind
	mapPlan            [][]string
	expeditionPosition Expedition
}
type TimeLine struct {
	timeFrames []TimeFrame
}

func (timeLine *TimeLine) addTimeFrame(timeFrame TimeFrame) {
	timeLine.timeFrames = append(timeLine.timeFrames, timeFrame)
}
func (timeLine *TimeLine) removeLastTimeFrame() {
	timeLine.timeFrames = timeLine.timeFrames[:len(timeLine.timeFrames)-1]
}

type Expedition struct {
	position Position
}

func (expedition *Expedition) getPossiblePositions(newMapPlan [][]string) []Position {
	nearPositions := make([]Position, 0)
	nearPositions = append(nearPositions, Position{expedition.position.x, expedition.position.y})
	nearPositions = append(nearPositions, Position{expedition.position.x, expedition.position.y + 1})
	nearPositions = append(nearPositions, Position{expedition.position.x + 1, expedition.position.y})
	nearPositions = append(nearPositions, Position{expedition.position.x, expedition.position.y - 1})
	nearPositions = append(nearPositions, Position{expedition.position.x - 1, expedition.position.y})

	possiblePositions := make([]Position, 0)
	for _, position := range nearPositions {
		if position.x >= 0 && position.x < len(newMapPlan) && position.y >= 0 && position.y < len(newMapPlan[0]) && newMapPlan[position.x][position.y] == "." {
			possiblePositions = append(possiblePositions, position)
		}
	}

	return possiblePositions
}
func (expedition *Expedition) updatePosition(position Position) {
	expedition.position = position
}

var (
	//go:embed input.txt
	input string

	ALLTimeLines   []TimeLine
	mapPlanModel   [][]string
	targetPosition Position
)

func init() {
	input := strings.TrimSuffix(input, "\n")
	content := strings.Split(input, "\n")

	/// create mapPlanModel
	for _, line := range content {
		mapLine := make([]string, 0)
		for _, char := range line {
			if char == '#' {
				mapLine = append(mapLine, "#")
			} else {
				mapLine = append(mapLine, ".")
			}
		}
		mapPlanModel = append(mapPlanModel, mapLine)
	}
	// mapPlanPrinter(mapPlanModel) // prints E in the first position because theExpedition.position is not initialized yet
	////////////////////////

	yend := 0
	for id, elem := range mapPlanModel[len(mapPlanModel)-1] {
		if elem == "." {
			yend = id
			break
		}
	}
	targetPosition = Position{len(mapPlanModel) - 1, yend}
	fmt.Print(targetPosition)

	ALLTimeLines = make([]TimeLine, 0)
}

func nextRound(timeLine TimeLine, MinuteCounter int) int {

	if MinuteCounter == 0 {
		// add to TimeFrames
		ALLTimeLines = append(ALLTimeLines, timeLine)
		return -1

	} else {

		lastFrame := timeLine.timeFrames[len(timeLine.timeFrames)-1]

		// create new windPositions
		newWindPositions := make(map[Position][]Wind)
		for _, winds := range lastFrame.windPositions {
			for _, wind := range winds {
				if wind.detectObstacle(lastFrame.mapPlan) {
					switch wind.direction {
					case 0:
						newWindPositions[Position{wind.position.x, 1}] = append(newWindPositions[Position{wind.position.x, 1}], Wind{Position{wind.position.x, 1}, wind.direction})
					case 1:
						newWindPositions[Position{1, wind.position.y}] = append(newWindPositions[Position{1, wind.position.y}], Wind{Position{1, wind.position.y}, wind.direction})
					case 2:
						newWindPositions[Position{wind.position.x, len(lastFrame.mapPlan[0]) - 2}] = append(newWindPositions[Position{wind.position.x, len(lastFrame.mapPlan[0]) - 2}], Wind{Position{wind.position.x, len(lastFrame.mapPlan[0]) - 2}, wind.direction})
					case 3:
						newWindPositions[Position{len(lastFrame.mapPlan) - 2, wind.position.y}] = append(newWindPositions[Position{len(lastFrame.mapPlan) - 2, wind.position.y}], Wind{Position{len(lastFrame.mapPlan) - 2, wind.position.y}, wind.direction})
					}
				} else {
					switch wind.direction {
					case 0:
						newWindPositions[Position{wind.position.x, wind.position.y + 1}] = append(newWindPositions[Position{wind.position.x, wind.position.y + 1}], Wind{Position{wind.position.x, wind.position.y + 1}, wind.direction})
					case 1:
						newWindPositions[Position{wind.position.x + 1, wind.position.y}] = append(newWindPositions[Position{wind.position.x + 1, wind.position.y}], Wind{Position{wind.position.x + 1, wind.position.y}, wind.direction})
					case 2:
						newWindPositions[Position{wind.position.x, wind.position.y - 1}] = append(newWindPositions[Position{wind.position.x, wind.position.y - 1}], Wind{Position{wind.position.x, wind.position.y - 1}, wind.direction})
					case 3:
						newWindPositions[Position{wind.position.x - 1, wind.position.y}] = append(newWindPositions[Position{wind.position.x - 1, wind.position.y}], Wind{Position{wind.position.x - 1, wind.position.y}, wind.direction})
					}
				}
			}
		}
		// create new mapPlan
		newMapPlan := make([][]string, len(mapPlanModel))
		newMapPlan = mapPlanCreator(newWindPositions, newMapPlan)
		// mapWindPrinter(newMapPlan)
		// time.Sleep(time.Second / 2)

		// analyse new ExpeditionPosition
		newExpeditionPosition := lastFrame.expeditionPosition // still the last here !!
		freePositions := newExpeditionPosition.getPossiblePositions(newMapPlan)

		//duplicating the timeLine
		newTimeLine := TimeLine{timeLine.timeFrames}

		if len(freePositions) == 0 {
			fmt.Println("no more free positions(number ", len(ALLTimeLines)-1, "in ALLTimeLines)")
			newTimeLine.addTimeFrame(TimeFrame{newWindPositions, newMapPlan, newExpeditionPosition})
			ALLTimeLines = append(ALLTimeLines, newTimeLine)
			return -2
		}

		var nextSteps []int
		for _, freePosition := range freePositions {
			newTimeLine.addTimeFrame(TimeFrame{newWindPositions, newMapPlan, newExpeditionPosition})
			if freePosition.x == targetPosition.x && freePosition.y == targetPosition.y {
				ALLTimeLines = append(ALLTimeLines, newTimeLine)
				fmt.Print("found it ! with: ", len(newTimeLine.timeFrames), " steps (number ", len(ALLTimeLines)-1, "in ALLTimeLines)")
				return len(newTimeLine.timeFrames)
			} else {
				nextSteps = append(nextSteps, nextRound(newTimeLine, MinuteCounter-1))
			}
			newTimeLine.removeLastTimeFrame()
		}
		return max(0, minArray(nextSteps))
	}
}

func timeLineVizualizer(timeLine TimeLine) {

	mapPlanPrinter(mapPlanModel, timeLine.timeFrames[0].expeditionPosition)
	time.Sleep(1 * time.Second)
	for _, timeFrame := range timeLine.timeFrames {
		mapWindPrinter(timeFrame.mapPlan)
		time.Sleep(time.Second / 2)
		mapPlanPrinter(timeFrame.mapPlan, timeFrame.expeditionPosition)
		time.Sleep(1 * time.Second)
	}
}

func part1() {

	/////// create initial TimeFrame
	/// create initial WindMap
	windPositions := make(map[Position][]Wind)
	content := strings.Split(input, "\n")
	for x, line := range content {
		for y, char := range line {
			switch char {
			case '>':
				windPositions[Position{x, y}] = append(windPositions[Position{x, y}], Wind{Position{x, y}, 0})
			case 'v':
				windPositions[Position{x, y}] = append(windPositions[Position{x, y}], Wind{Position{x, y}, 1})
			case '<':
				windPositions[Position{x, y}] = append(windPositions[Position{x, y}], Wind{Position{x, y}, 2})
			case '^':
				windPositions[Position{x, y}] = append(windPositions[Position{x, y}], Wind{Position{x, y}, 3})
			}
		}
	}
	/// create initial mapPlan
	mapPlan := make([][]string, len(mapPlanModel))
	for i := range mapPlan {
		mapPlan[i] = make([]string, len(mapPlanModel[i]))
		copy(mapPlan[i], mapPlanModel[i])
	}
	for position, winds := range windPositions {
		switch winds[0].direction {
		case 0:
			mapPlan[position.x][position.y] = ">"
		case 1:
			mapPlan[position.x][position.y] = "v"
		case 2:
			mapPlan[position.x][position.y] = "<"
		case 3:
			mapPlan[position.x][position.y] = "^"
		}
	}
	/// create initial expeditionPosition
	var InitialExpeditionPosition Expedition
	ystart := 0
	for id, elem := range mapPlanModel[0] {
		if elem == "." {
			ystart = id
			break
		}
	}
	InitialExpeditionPosition = Expedition{Position{0, ystart}}
	/// finalize initial timeFrame
	initialTimeFrame := TimeFrame{windPositions, mapPlan, InitialExpeditionPosition}
	/// create timeLine // which will be modified each Round
	var timeLine TimeLine
	timeLine.addTimeFrame(initialTimeFrame)
	////////////////////////

	mapPlan = mapPlanCreator(windPositions, mapPlan)
	mapPlanPrinter(mapPlan, InitialExpeditionPosition)

	// timeLineVizualizer(timeLine)

	/////// calculating all possible timeLines :
	/////// nextRound : recursive function that adds a timeLine when the Expedition is blocked or achieves last destination

	MinuteCounter := 500
	fmt.Println(nextRound(timeLine, MinuteCounter), len(ALLTimeLines))

}

func main() {

	start := time.Now()
	fmt.Print("\npart1: ")
	part1()
	fmt.Println(time.Since(start))
	timeLineVizualizer(ALLTimeLines[0])

}
