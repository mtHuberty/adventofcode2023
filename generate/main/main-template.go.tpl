package main

import (
	"fmt"
	"os"
    "log"

    {{- range .Days }}
    "github.com/mtHuberty/adventofcode2023/days/{{ . }}"
    {{- end }}
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatal("Not enough arguments provided")
	}

    switch args[0] {
    {{- range .Days }}
    case "{{ . }}":
        switch args[1] {
        case "one":
            fmt.Println({{ . }}.SolvePartOne())
        case "two":
            fmt.Println({{ . }}.SolvePartTwo())
        default:
            log.Fatal("Unknown second argument")
        }
    {{- end }}
    default:
        log.Fatal("Unknown first argument")
    }
}
