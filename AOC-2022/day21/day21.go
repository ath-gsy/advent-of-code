package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
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

	waitedList = make(map[string]WaitingMonkeys) // not used yet
	//key = ; value =  ///to define

)

func init() {
	input := strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		monkeyInfos := strings.Split(line, ": ")

		monkeyName := monkeyInfos[0]
		orderTable = append(orderTable, monkeyName) // pour garder l'ordre

		////INIT MONKEYSMAP: met un number si possible, sinon set up l'attente avec -1 pour les numéros absents
		number, err := strconv.Atoi(monkeyInfos[1])
		if err == nil {
			monkeysMap[monkeyName] = Instruction{number: number}
		} else { // BIG assumption on the fact that monkeys don't yell negative numbers //
			infos := strings.Split(monkeyInfos[1], " ")
			monkeysMap[monkeyName] = Instruction{number: -1, monkey1: infos[0], monkey1Number: -1, operation: infos[1], monkey2: infos[2], monkey2Number: -1}
		}
		//////////////////////////////////////////////////////////////////////////////
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

func main() {
	fmt.Println(monkeysMap)
	fmt.Println(orderTable)

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
					result = monkey.monkey1Number / monkey.monkey2Number
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

	fmt.Println(orderTable)
	fmt.Println(monkeysMap)
	fmt.Println(monkeysMap["root"].number)

}
