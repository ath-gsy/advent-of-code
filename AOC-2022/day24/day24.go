package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

func init() {
	input := strings.TrimSuffix(input, "\n")
	content := strings.Split(input, "\n\n")

	fmt.Print(content)

}

func main() {
}
