// 1er code en go (je ne pense pas pouvoir finir aujourd'hui)

package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type MixValue struct {
	mixingOrder int
	value       int
}

var (
	//go:embed input.txt
	input string

	globalMap       = make(map[int]MixValue)
	totalLength int = 0
)

func init() {
	input := strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	for originalOrder, value := range lines {

		var value, _ = strconv.Atoi(value)
		globalMap[originalOrder] = MixValue{originalOrder, value}
	}

	totalLength = len(lines)
	// fmt.Println(totalLength)
}

func main() {
	// fmt.Println(globalMap)

	originalOrder := 0
	for originalOrder < totalLength {
		mixOrder(originalOrder)
		originalOrder++
	}

	// fmt.Println(globalMap)

	var finalOrder []int

	for i := 0; i < totalLength; i++ {
		finalOrder = append(finalOrder, globalMap[mapkey(i)].value)
	}

	var index0 int
	for id, v := range finalOrder {
		if v == 0 {
			index0 = id
		}
	}
	fmt.Println(finalOrder)
	fmt.Println(index0)

	var sumCoordinates int
	var counter int
	currId := index0
	for counter != 1000 {
		counter++
		currId++
		if currId == totalLength {
			currId = 0
		}
	}
	sumCoordinates += finalOrder[currId]
	for counter != 2000 {
		counter++
		currId++
		if currId == totalLength {
			currId = 0
		}
	}
	sumCoordinates += finalOrder[currId]
	for counter != 3000 {
		counter++
		currId++
		if currId == totalLength {
			currId = 0
		}
	}
	sumCoordinates += finalOrder[currId]

	fmt.Println(sumCoordinates)
}

func mapkey(id int) (key int) { // cherche la clé de l'élément qui est classé à l'indice i ("", ok bool" : en sortie, enlevé)
	for k, v := range globalMap {
		if v.mixingOrder == id {
			key = k
			// ok = true
			return
		}
	}
	return
}

func mixOrder(originalOrder int) {
	var displacementValue int = globalMap[originalOrder].value
	var mixOrder = globalMap[originalOrder].mixingOrder

	var newOrder = (mixOrder + displacementValue) % totalLength

	if newOrder < 0 {
		newOrder = totalLength + newOrder
	}

	newEntry := globalMap[originalOrder]
	newEntry.mixingOrder = newOrder

	var tempValue MixValue
	var tempMapKey int

	if displacementValue > 0 { //vers le bas ou...
		if newOrder < mixOrder { //cas particulier overlap
			for i := mixOrder; i < totalLength-1; i++ {
				tempMapKey = mapkey(i + 1)
				tempValue = globalMap[tempMapKey]
				tempValue.mixingOrder = i
				globalMap[tempMapKey] = tempValue
			}
			//cas particulier de length à 0
			tempMapKey = mapkey(0)
			tempValue = globalMap[tempMapKey]
			tempValue.mixingOrder = totalLength - 1
			globalMap[tempMapKey] = tempValue
			//////////////////////////////
			for i := 0; i < newOrder; i++ {
				tempMapKey = mapkey(i + 1)
				tempValue = globalMap[tempMapKey]
				tempValue.mixingOrder = i
				globalMap[tempMapKey] = tempValue
			}
		} else {
			for i := mixOrder; i < newOrder; i++ {
				tempMapKey = mapkey(i + 1)
				tempValue = globalMap[tempMapKey]
				tempValue.mixingOrder = i
				globalMap[tempMapKey] = tempValue

			}
		}
	} else if displacementValue < 0 { //...vers le haut
		if newOrder > mixOrder { //cas particulier overlap
			for i := mixOrder; i > 0; i-- {
				tempMapKey = mapkey(i - 1)
				tempValue = globalMap[tempMapKey]
				tempValue.mixingOrder = i
				globalMap[tempMapKey] = tempValue
			}
			//cas particulier de 0 à length
			tempMapKey = mapkey(totalLength - 1)
			tempValue = globalMap[tempMapKey]
			tempValue.mixingOrder = 0
			globalMap[tempMapKey] = tempValue
			//////////////////////////////
			for i := totalLength - 1; i > newOrder; i-- {
				tempMapKey = mapkey(i - 1)
				tempValue = globalMap[tempMapKey]
				tempValue.mixingOrder = i
				globalMap[tempMapKey] = tempValue
			}
		} else {
			for i := mixOrder; i > newOrder; i-- {
				tempMapKey = mapkey(i - 1)
				tempValue = globalMap[tempMapKey]
				tempValue.mixingOrder = i
				globalMap[tempMapKey] = tempValue
			}
		}
	}

	globalMap[originalOrder] = newEntry

}
