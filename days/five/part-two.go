package five

import (
	"regexp"
	"sort"
	"strconv"

	"github.com/mtHuberty/adventofcode2023/util"
)

var linesTwo []string

func readFileTwo() {
	scanner := util.NewScanner("./days/five/input.txt")

	defer scanner.Close()

	for scanner.Scan() {
		line := scanner.Text()
		linesTwo = append(linesTwo, line)
	}
}

func SolvePartTwo() string {
	readFileTwo()
	var sourceNums []int

	mapKey := ""

	cmap := make(conversionMap)

	for i, line := range linesOne {
		if i == 0 {
			sourceNums = handleFirstLine(line)
			continue
		}

		// if line is empty, skip
		if line == "" {
			continue
		}

		// if line starts with "(something)-to-(something else) map:"
		// add a new key to our map with the key being "(something)-to-(something else)"
		r := regexp.MustCompile(`([a-z]+-to-[a-z]+) map:`)

		if r.MatchString(line) {
			matches := r.FindStringSubmatch(line)
			mapKey = matches[1]
			// fmt.Println("NOW HANDLING", mapKey)
			cmap[mapKey] = []conversionInfo{}
			continue
		}

		// if line matches "number1 number2 number3"
		// add a new conversionInfo to our map with the source being number1, dest being number2, and rangeVal being number3
		r = regexp.MustCompile(`([0-9]+) ([0-9]+) ([0-9]+)`)
		if r.MatchString(line) {
			matches := r.FindStringSubmatch(line)
			destStart, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			sourceStart, err := strconv.Atoi(matches[2])
			if err != nil {
				panic(err)
			}
			rangeVal, err := strconv.Atoi(matches[3])
			if err != nil {
				panic(err)
			}

			cmap[mapKey] = append(cmap[mapKey], conversionInfo{destStart, sourceStart, rangeVal})
			continue
		}
	}

	results := []int{}
	for _, sourceNum := range sourceNums {
		results = append(results, convertSourceToDest(sourceNum, "seed", &cmap))
	}

	// ignore empty lines
	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})

	// fmt.Println(results)

	return strconv.Itoa(results[0])
}
