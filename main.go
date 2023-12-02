package main

import (
	"fmt"
	"os"

	"github.com/mtHuberty/adventofcode2023/one"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}

	switch args[0] {
	case "one":
		switch args[1] {
		case "one":
			fmt.Println(one.SolvePartOne())
		case "two":
			fmt.Println(one.SolvePartTwo())
		}
	default:
		fmt.Println("Unknown argument")
		os.Exit(1)
	}
}
