package five

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/mtHuberty/adventofcode2023/util"
	"golang.org/x/sync/errgroup"
)

// BROKEN IN CURRENT STATE
// I regret choosing recursion for part one. But also I think part 2 (non brute force)
// requires a different approach altogether. Might need to convert "ranges" instead of
// individual seed numbers, and create new "ranges" when needed.

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
	fmt.Println("reading file...")
	readFileTwo()
	fmt.Println("done reading file.")
	var sourceNums []int

	mapKey := ""

	cmap := make(conversionMap)

	for i, line := range linesTwo {
		if i == 0 {
			fmt.Println("handling first line")
			sourceNums = handleFirstLineAsRanges(line)
			continue
		}

		// if line is empty, skip
		if line == "" {
			fmt.Println("handling empty line")
			continue
		}

		// if line starts with "(something)-to-(something else) map:"
		// add a new key to our map with the key being "(something)-to-(something else)"
		r := regexp.MustCompile(`([a-z]+-to-[a-z]+) map:`)
		if r.MatchString(line) {
			fmt.Println("handling conversion title line")
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
			fmt.Println("handling conversion line")
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

	results := safeIntSlice{
		mu:    sync.Mutex{},
		slice: []int{},
	}

	fmt.Println("starting conversions")

	g, _ := errgroup.WithContext(context.Background())
	for _, sourceNum := range sourceNums {
		num := sourceNum // pin the num
		g.Go(func() error {
			asyncConvertSourceToDest(num, "seed", &cmap, &results)
			return nil
		})
	}
	g.Wait()
	fmt.Println("done with conversions")

	// ignore empty lines
	sort.Slice(results.slice, func(i, j int) bool {
		return results.slice[i] < results.slice[j]
	})

	// fmt.Println(results)

	return strconv.Itoa(results.slice[0])
}

type safeIntSlice struct {
	slice []int
	mu    sync.Mutex
}

func (s *safeIntSlice) append(vals []int) {
	s.mu.Lock()
	s.slice = append(s.slice, vals...)
	s.mu.Unlock()
}

// First line starts "seeds: number1 number2...."
// Same as handleFirstLine, but recognizes that the
// numbers come in pairs as startingNum and range.
func handleFirstLineAsRanges(line string) []int {
	seedNumStrings := strings.Split(line, " ")[1:]

	// Make sure we have an even number so we don't
	// blow up down below accessing [i+1]
	if len(seedNumStrings)%2 != 0 {
		panic("seed num count not an even number")
	}

	safeSeedNumSlice := safeIntSlice{
		mu:    sync.Mutex{},
		slice: make([]int, len(seedNumStrings)),
	}

	g, _ := errgroup.WithContext(context.Background())

	for i := range seedNumStrings {
		fmt.Println("working on i", i)
		if i%2 != 0 {
			continue
		}
		startNumStr := seedNumStrings[i]
		rangeLenStr := seedNumStrings[i+1]

		startNum, err := strconv.Atoi(startNumStr)
		if err != nil {
			panic(err)
		}
		rangeLen, err := strconv.Atoi(rangeLenStr)
		if err != nil {
			panic(err)
		}

		g.Go(func() error {
			asyncAddSeedNumsFromRange(startNum, rangeLen, &safeSeedNumSlice)
			return nil
		})
	}

	g.Wait()
	fmt.Println("done with async processing")

	return safeSeedNumSlice.slice
}

func asyncAddSeedNumsFromRange(startNum int, rangeLen int, s *safeIntSlice) {
	nums := []int{}
	for j := 0; j < rangeLen; j++ {
		nums = append(nums, startNum+j)
	}
	s.append(nums)
}

// Recursively converts sourceNum of type sourceType through to the final type
func asyncConvertSourceToDest(sourceNum int, sourceType string, cmap *conversionMap, s *safeIntSlice) int {
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
					return asyncConvertSourceToDest(destNum, destType, cmap, s)
				}
			}

			// fmt.Printf("did not find sourceNum %d in data %v\n", sourceNum, conversionInfos)
			// fmt.Printf("source num %d is dest num %d\n", sourceNum, sourceNum)

			return asyncConvertSourceToDest(sourceNum, destType, cmap, s)
		}
	}

	return sourceNum
}
