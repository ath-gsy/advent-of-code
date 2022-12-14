package other

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func mapPlanPrinter(mapPlan [][]string) {
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
	windPositions map[Position][]Wind
	mapPlan       [][]string
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

	timeFrames   []TimeFrame
	mapPlanModel [][]string

	theExpeditionPosition Expedition
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

	/////// create initial TimeFrame
	/// create initial WindMap
	windPositions := make(map[Position][]Wind)
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
		///
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
		// mapPlanPrinter(mapPlan)
		// mapWindPrinter(mapPlan) has been created after
		///
		timeFrames = append(timeFrames, TimeFrame{windPositions, mapPlan})
		////////////////////////

		// create expedition
		ystart := 0
		for id, elem := range mapPlanModel[0] {
			if elem == "." {
				ystart = id
				break
			}
		}
		theExpeditionPosition = Expedition{Position{0, ystart}}
		mapPlanPrinter(mapPlan)
		////////////////////////
	}
}

func part1() {
	time.Sleep(1 * time.Second)

	MINUTE_COUNTER := 0
	for MINUTE_COUNTER < 10 { // a remplacer par une condition sur le déplacement de l'expédition

		/////// create new TimeFrame

		// create new windPositions
		newWindPositions := make(map[Position][]Wind)
		for _, winds := range timeFrames[len(timeFrames)-1].windPositions {
			for _, wind := range winds {
				if wind.detectObstacle(timeFrames[len(timeFrames)-1].mapPlan) {
					switch wind.direction {
					case 0:
						newWindPositions[Position{wind.position.x, 1}] = append(newWindPositions[Position{wind.position.x, 1}], Wind{Position{wind.position.x, 1}, wind.direction})
					case 1:
						newWindPositions[Position{1, wind.position.y}] = append(newWindPositions[Position{1, wind.position.y}], Wind{Position{1, wind.position.y}, wind.direction})
					case 2:
						newWindPositions[Position{wind.position.x, len(timeFrames[len(timeFrames)-1].mapPlan[0]) - 2}] = append(newWindPositions[Position{wind.position.x, len(timeFrames[len(timeFrames)-1].mapPlan[0]) - 2}], Wind{Position{wind.position.x, len(timeFrames[len(timeFrames)-1].mapPlan[0]) - 2}, wind.direction})
					case 3:
						newWindPositions[Position{len(timeFrames[len(timeFrames)-1].mapPlan) - 2, wind.position.y}] = append(newWindPositions[Position{len(timeFrames[len(timeFrames)-1].mapPlan) - 2, wind.position.y}], Wind{Position{len(timeFrames[len(timeFrames)-1].mapPlan) - 2, wind.position.y}, wind.direction})
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
		mapWindPrinter(newMapPlan)
		time.Sleep(time.Second / 2)
		// add to TimeFrames
		timeFrames = append(timeFrames, TimeFrame{newWindPositions, newMapPlan})
		////////////////////////

		/////// expedition analyse

		freePositions := theExpeditionPosition.getPossiblePositions(newMapPlan)

		fmt.Println(freePositions)

		if len(freePositions) > 1 {
			theExpeditionPosition.updatePosition(freePositions[1])
		} else if len(freePositions) == 1 {
			theExpeditionPosition.updatePosition(freePositions[0])
		} else {
			fmt.Println("no more free positions")
			break
		}
		mapPlanPrinter(newMapPlan)
		////////////////////////

		//counter
		MINUTE_COUNTER++
		time.Sleep(2 * time.Second)
	}
}

func main() {

	start := time.Now()
	fmt.Print("\npart1: ")
	part1()
	fmt.Println(time.Since(start))

}
