package main

import (
	"fmt"
	"os"
    "log"
    "github.com/mtHuberty/adventofcode2023/days/three"
    "github.com/mtHuberty/adventofcode2023/days/one"
    "github.com/mtHuberty/adventofcode2023/days/two"
    "github.com/mtHuberty/adventofcode2023/days/four"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatal("Not enough arguments provided")
	}

    switch args[0] {
    case "three":
        switch args[1] {
        case "one":
            fmt.Println(three.SolvePartOne())
        case "two":
            fmt.Println(three.SolvePartTwo())
        default:
            log.Fatal("Unknown second argument")
        }
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
    case "four":
        switch args[1] {
        case "one":
            fmt.Println(four.SolvePartOne())
        case "two":
            fmt.Println(four.SolvePartTwo())
        default:
            log.Fatal("Unknown second argument")
        }
    default:
        log.Fatal("Unknown first argument")
    }
}
