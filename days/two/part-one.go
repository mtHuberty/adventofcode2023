// template.go.tpl

package two

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// only 12 red cubes, 13 green cubes, and 14 blue cubes?
var cubeCountMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func SolvePartOne() string {
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
		gameNumSubstring := substrings[0]                          // "Game 1"
		handfulsSubstring := substrings[1]                         // "12 blue, 11 green, 3 red; 6 blue, 5 green, 7 red; 5 red, 11 blue; 2 blue, 8 green"
		handfulSubstrings := strings.Split(handfulsSubstring, ";") // [ "12 blue, 11 green, 3 red", " 6 blue, 5 green, 7 red", " 5 red, 11 blue", " 2 blue, 8 green" ]

		gameNum, err := strconv.Atoi(gameNumSubstring[5:]) // 1
		if err != nil {
			log.Fatal("failed to convert gamenum to int:", err)
		}

		possible := true

		for _, handful := range handfulSubstrings {
			handful = strings.TrimSpace(handful) // "12 blue, 11 green, 3 red"

			colorResults := strings.Split(handful, ",") // [ "12 blue", " 11 green", " 3 red" ]

			for _, colorResult := range colorResults {
				colorResult = strings.TrimSpace(colorResult) // "12 blue"

				colorResultSplit := strings.Split(colorResult, " ") // [ "12", "blue" ]

				cubeCountStr := colorResultSplit[0]
				cubeColor := colorResultSplit[1]

				cubeCount, err := strconv.Atoi(cubeCountStr)
				if err != nil {
					log.Fatalf("error converting cubeCountStr %s to int: %s", cubeCountStr, err)
				}

				if cubeCount > cubeCountMap[cubeColor] {
					possible = false
					break
				}
			}
		}

		if possible {
			total += gameNum
		}
	}

	return strconv.Itoa(total)
}
