package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Instruction struct {
	number        int
	monkey1       string
	monkey1Number int
	operation     string
	monkey2       string
	monkey2Number int
}

type WaitingMonkeys struct {
	monkeys []string
}

var (
	//go:embed input.txt
	input string

	orderTable []string
	monkeysMap = make(map[string]Instruction)

	ORIGINALorderTable []string                       //pour la partie2
	ORIGINALmonkeysMap = make(map[string]Instruction) //pour la partie2

	tmpOrderTable []string                       //pour la partie2
	tmpMonkeysMap = make(map[string]Instruction) //pour la partie2

)

func init() {
	input := strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		monkeyInfos := strings.Split(line, ": ")

		monkeyName := monkeyInfos[0]
		orderTable = append(orderTable, monkeyName) // pour garder l'ordre
		ORIGINALorderTable = append(ORIGINALorderTable, monkeyName)

		////INIT MONKEYSMAP: met un number si possible, sinon set up l'attente avec -1 pour les numéros absents
		number, err := strconv.Atoi(monkeyInfos[1])
		if err == nil {
			monkeysMap[monkeyName] = Instruction{number: number}
			ORIGINALmonkeysMap[monkeyName] = Instruction{number: number} //pour la partie2
		} else { // BIG assumption on the fact that monkeys don't yell negative numbers //
			infos := strings.Split(monkeyInfos[1], " ")
			monkeysMap[monkeyName] = Instruction{number: -1, monkey1: infos[0], monkey1Number: -1, operation: infos[1], monkey2: infos[2], monkey2Number: -1}
			ORIGINALmonkeysMap[monkeyName] = Instruction{number: -1, monkey1: infos[0], monkey1Number: -1, operation: infos[1], monkey2: infos[2], monkey2Number: -1} //pour la partie2
		}
		//////////////////////////////////////////////////////////////////////////////
	}

	for index, element := range ORIGINALmonkeysMap {
		tmpMonkeysMap[index] = element
	}

}

func receiveNumber(monkeyYeller string, number int) { // receiveNumber from monkey "monkeyName" by everyone waiting it

	for monkey, instructions := range monkeysMap {
		if instructions.number == -1 {
			if instructions.monkey1 == monkeyYeller {
				instructions.monkey1Number = number
			} else if instructions.monkey2 == monkeyYeller {
				instructions.monkey2Number = number
			}
			monkeysMap[monkey] = instructions
		}

	}
}

func receiveNumber2(monkeyYeller string, number int) { // receiveNumber from monkey "monkeyName" by everyone waiting it

	for monkey, instructions := range tmpMonkeysMap {
		if instructions.number == -1 {
			if instructions.monkey1 == monkeyYeller {
				instructions.monkey1Number = number
			} else if instructions.monkey2 == monkeyYeller {
				instructions.monkey2Number = number
			}
			tmpMonkeysMap[monkey] = instructions
		}

	}
}

func part1() {
	// fmt.Println(monkeysMap)
	// fmt.Println(orderTable)

	for monkeysMap["root"].number == -1 {
		var newOrderTable []string
		for _, monkeyName := range orderTable {
			monkey := monkeysMap[monkeyName] //copy of the map element
			if monkey.number != -1 {

				receiveNumber(monkeyName, monkey.number)

			} else if monkey.monkey1Number != -1 && monkey.monkey2Number != -1 {
				var result int
				if monkey.operation == "+" {
					result = monkey.monkey1Number + monkey.monkey2Number
				} else if monkey.operation == "*" {
					result = monkey.monkey1Number * monkey.monkey2Number
				} else if monkey.operation == "-" {
					result = monkey.monkey1Number - monkey.monkey2Number
				} else if monkey.operation == "/" {
					if monkey.monkey2Number != 0 {
						result = monkey.monkey1Number / monkey.monkey2Number
					} else {
						result = 0
						fmt.Println("division by zero")
					}
				}
				monkey.number = result
				monkeysMap[monkeyName] = monkey
				receiveNumber(monkeyName, result)

			} else {
				newOrderTable = append(newOrderTable, monkeyName) // leave the ones who didn't give information yet in the orderTable
			}
			// monkeysMap[monkeyName] = monkey
		}
		orderTable = newOrderTable
	}

	// fmt.Println(orderTable)
	// fmt.Println(monkeysMap)
	fmt.Println(monkeysMap["root"].number)
	fmt.Println(monkeysMap["root"])

}

// func oneTry(humnNumber) {

// }

func part2() {
	// fmt.Println(monkeysMap)
	// fmt.Println(orderTable)

	// tmpOrderTable := make([]string, len(ORIGINALorderTable))
	copy(tmpOrderTable, ORIGINALorderTable)

	counter := 0

	// humnInfo := tmpMonkeysMap["humn"]
	// humnInfo.number = 0 + counter
	// tmpMonkeysMap["humn"] = humnInfo

	for !((tmpMonkeysMap["root"].monkey1Number != -1 || tmpMonkeysMap["root"].monkey2Number != -1) && tmpMonkeysMap["root"].monkey1Number == tmpMonkeysMap["root"].monkey2Number) {

		///restart map and ordertable
		tmpOrderTable := make([]string, len(ORIGINALorderTable))
		copy(tmpOrderTable, ORIGINALorderTable)
		for index, element := range ORIGINALmonkeysMap {
			tmpMonkeysMap[index] = element
		}
		humnInfo := tmpMonkeysMap["humn"]
		counter++
		humnInfo.number = 0 + counter

		tmpMonkeysMap["humn"] = humnInfo
		///

		// fmt.Println(tmpOrderTable)
		fmt.Println(tmpMonkeysMap["humn"])
		fmt.Println(tmpMonkeysMap["root"])
		// fmt.Println(tmpMonkeysMap)

		for tmpMonkeysMap["root"].number == -1 {
			// fmt.Println(tmpOrderTable)
			// fmt.Println(tmpMonkeysMap)
			var newOrderTable []string
			for _, monkeyName := range tmpOrderTable {
				monkey := tmpMonkeysMap[monkeyName] //copy of the map element
				if monkey.number != -1 {

					receiveNumber2(monkeyName, monkey.number)

				} else if monkey.monkey1Number != -1 && monkey.monkey2Number != -1 {
					var result int
					if monkey.operation == "+" {
						result = monkey.monkey1Number + monkey.monkey2Number
					} else if monkey.operation == "*" {
						result = monkey.monkey1Number * monkey.monkey2Number
					} else if monkey.operation == "-" {
						result = monkey.monkey1Number - monkey.monkey2Number
					} else if monkey.operation == "/" {
						if monkey.monkey2Number != 0 {
							result = monkey.monkey1Number / monkey.monkey2Number
						} else {
							result = 0
							fmt.Println("division by zero")
						}
					}
					monkey.number = result
					tmpMonkeysMap[monkeyName] = monkey
					receiveNumber2(monkeyName, result)

				} else {
					fmt.Println(newOrderTable)
					newOrderTable = append(newOrderTable, monkeyName) // leave the ones who didn't give information yet in the orderTable
				}
				// monkeysMap[monkeyName] = monkey
			}
			tmpOrderTable = newOrderTable

		}
		fmt.Println(tmpMonkeysMap["root"])
		fmt.Println(tmpMonkeysMap["root"].monkey1Number, tmpMonkeysMap["root"].monkey2Number)

	}

	// fmt.Println(orderTable)
	// fmt.Println(monkeysMap)

}

func main() {

	start := time.Now()
	fmt.Print("part1: ")
	part1()
	fmt.Println(time.Since(start))

	// start := time.Now()
	// fmt.Print("part2: ")
	// part2()
	// fmt.Println(time.Since(start))

}
