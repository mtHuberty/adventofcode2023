// template.go.tpl

package two

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func SolvePartTwo() string {
	file, err := os.Open("./days/two/input.txt")
	if err != nil {
		log.Fatal("could not open file:", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Game 1: 12 blue, 11 green, 3 red; 6 blue, 5 green, 7 red; 5 red, 11 blue; 2 blue, 8 green
		substrings := strings.Split(line, ":")
		handfulsSubstring := substrings[1]                         // "12 blue, 11 green, 3 red; 6 blue, 5 green, 7 red; 5 red, 11 blue; 2 blue, 8 green"
		handfulSubstrings := strings.Split(handfulsSubstring, ";") // [ "12 blue, 11 green, 3 red", " 6 blue, 5 green, 7 red", " 5 red, 11 blue", " 2 blue, 8 green" ]

		maxRedCubes := 0
		maxGreenCubes := 0
		maxBlueCubes := 0

		for _, handful := range handfulSubstrings {
			handful = strings.TrimSpace(handful) // "12 blue, 11 green, 3 red"

			colorResults := strings.Split(handful, ",") // [ "12 blue", " 11 green", " 3 red" ]

			for _, colorResult := range colorResults {
				colorResult = strings.TrimSpace(colorResult) // "12 blue"

				colorResultSplit := strings.Split(colorResult, " ") // [ "12", "blue" ]

				cubeCountStr := colorResultSplit[0] // "12"
				cubeColor := colorResultSplit[1]    // "blue"

				cubeCount, err := strconv.Atoi(cubeCountStr) // 12
				if err != nil {
					log.Fatalf("error converting cubeCountStr %s to int: %s", cubeCountStr, err)
				}

				switch cubeColor {
				case "red":
					if cubeCount > maxRedCubes {
						maxRedCubes = cubeCount
					}
				case "green":
					if cubeCount > maxGreenCubes {
						maxGreenCubes = cubeCount
					}
				case "blue":
					if cubeCount > maxBlueCubes {
						maxBlueCubes = cubeCount
					}
				}
			}
		}

		power := maxRedCubes * maxGreenCubes * maxBlueCubes

		total += power
	}

	return strconv.Itoa(total)
}
