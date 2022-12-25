package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

var (
	//go:embed input.txt
	input string

	fuelList []string
)

func init() {
	input := strings.TrimSuffix(input, "\n")
	content := strings.Split(input, "\n")

	fuelList = append(fuelList, content...)
}

func SNAFUtoIntConverter(SNAFUstring string) int {
	numbers := strings.Split(SNAFUstring, "")
	var SNAFUint int

	var powerOf5 = 0.0
	for id := len(numbers) - 1; id >= 0; id-- {
		switch numbers[id] {
		case "2":
			SNAFUint += 2 * int(math.Pow(5, powerOf5))
		case "1":
			SNAFUint += 1 * int(math.Pow(5, powerOf5))
		case "0":
			SNAFUint += 0 * int(math.Pow(5, powerOf5))
		case "-":
			SNAFUint += -1 * int(math.Pow(5, powerOf5))
		case "=":
			SNAFUint += -2 * int(math.Pow(5, powerOf5))
		}
		powerOf5++
	}
	return SNAFUint
}

func intToSNAFUconverter(number int) string {
	var reverseOrderString string

	SNAFUremainder := 0
	for number != 0 {
		reste := number%5 + 1*SNAFUremainder
		SNAFUremainder = 0
		switch reste {
		case 2:
			reverseOrderString += "2"
		case 1:
			reverseOrderString += "1"
		case 0:
			reverseOrderString += "0"
		case 3:
			reverseOrderString += "="
			SNAFUremainder = 1
		case 4:
			reverseOrderString += "-"
			SNAFUremainder = 1
		case 5:
			reverseOrderString += "0"
			SNAFUremainder = 1
		}
		number /= 5
	}
	if SNAFUremainder == 1 {
		reverseOrderString += "1"
	}

	var SNAFUstring string
	for id := len(reverseOrderString) - 1; id >= 0; id-- {
		SNAFUstring += string(reverseOrderString[id])
	}

	return SNAFUstring
}

func part1() {

	totalFuel := 0

	for _, fuelNumber := range fuelList {
		// fmt.Println(fuelNumber)

		totalFuel += SNAFUtoIntConverter(fuelNumber)
	}

	fmt.Println(intToSNAFUconverter(totalFuel))
}

func main() {

	start := time.Now()
	fmt.Print("\npart1: ")
	part1()
	fmt.Println(time.Since(start))

}
