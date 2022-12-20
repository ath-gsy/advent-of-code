// 1er code en go (je ne pense pas pouvoir finir aujourd'hui)

package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	input := strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	fmt.Println(lines[0])

}
