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
	//go:embed exemple.txt
	input string

	globalMap   = make(map[int]MixValue)
	totalLength = 0
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
	fmt.Println(globalMap)

	originalOrder := 0
	for originalOrder < totalLength {
		mixOrder(originalOrder)
		originalOrder++
	}

	fmt.Println(globalMap)

}

func mapkey(id int) (key int) { //, ok bool en sortie enlevé
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
			for i := mixOrder; i > 1; i-- {
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
