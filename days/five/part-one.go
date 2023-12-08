package five

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/mtHuberty/adventofcode2023/util"
)

var linesOne []string

func readFileOne() {
	scanner := util.NewScanner("./days/five/input.txt")

	defer scanner.Close()

	for scanner.Scan() {

		line := scanner.Text()
		linesOne = append(linesOne, line)
	}
}

type conversionInfo struct {
	destStart   int
	sourceStart int
	rangeLen    int
}

type conversionMap map[string][]conversionInfo

func SolvePartOne() string {
	readFileOne()
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

// First line starts "seeds: number1 number2...."
func handleFirstLine(line string) []int {
	seedNumStrings := strings.Split(line, " ")[1:]

	seedNums := make([]int, len(seedNumStrings))

	for i, seedNumString := range seedNumStrings {
		s, err := strconv.Atoi(seedNumString)
		if err != nil {
			panic(err)
		}

		seedNums[i] = s
	}

	return seedNums
}

// Recursively converts sourceNum of type sourceType through to the final type
func convertSourceToDest(sourceNum int, sourceType string, cmap *conversionMap) int {
	fmt.Printf("converting sourceNum %d of type %s to dest", sourceNum, sourceType)
	if cmap == nil {
		panic("conversionMap can't be nil")
	}

	for mapKey, conversionInfos := range *cmap {
		if strings.HasPrefix(mapKey, sourceType) {
			destType := strings.Split(mapKey, "-")[2]

			// fmt.Printf("source type: %s to dest type: %s\n", sourceType, destType)

			for _, conversionInfo := range conversionInfos {
				if sourceNum >= conversionInfo.sourceStart && sourceNum <= conversionInfo.sourceStart+conversionInfo.rangeLen {
					destNum := (sourceNum - conversionInfo.sourceStart) + conversionInfo.destStart
					return convertSourceToDest(destNum, destType, cmap)
				}
			}

			// fmt.Printf("did not find sourceNum %d in data %v\n", sourceNum, conversionInfos)
			// fmt.Printf("source num %d is dest num %d\n", sourceNum, sourceNum)

			return convertSourceToDest(sourceNum, destType, cmap)
		}
	}

	return sourceNum
}
