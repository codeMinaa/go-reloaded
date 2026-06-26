package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("error reading input file:", err)
		return
	}
	lines := strings.Split(string(data), "\n")

	for i := range lines {
		lines[i] = processor(lines[i])
	}

	result := strings.Join(lines, "\n")

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("error writing output file:", err)
		return
	}

	fmt.Println("succeeded")
}
