package main

import (
	"fmt"
	"os"
    "log"
    "github.com/mtHuberty/adventofcode2023/days/one"
    "github.com/mtHuberty/adventofcode2023/days/two"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatal("Not enough arguments provided")
	}

    switch args[0] {
    case "one":
        switch args[1] {
        case "one":
            fmt.Println(one.SolvePartOne())
        case "two":
            fmt.Println(one.SolvePartTwo())
        default:
            log.Fatal("Unknown second argument")
        }
    case "two":
        switch args[1] {
        case "one":
            fmt.Println(two.SolvePartOne())
        case "two":
            fmt.Println(two.SolvePartTwo())
        default:
            log.Fatal("Unknown second argument")
        }
    default:
        log.Fatal("Unknown first argument")
    }
}
