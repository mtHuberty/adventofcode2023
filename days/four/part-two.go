package four

import (
	"strconv"
	"strings"

	"github.com/mtHuberty/adventofcode2023/util"
)

func SolvePartTwo() string {
	scanner := util.NewScanner("./days/four/input.txt")

	defer scanner.Close()

	total := 0

	cardWinMultiplier := map[int]int{}

	cardNum := 0

	for scanner.Scan() {
		cardNum++

		// fmt.Println("processing card", cardNum)

		if cardWinMultiplier[cardNum] == 0 {
			cardWinMultiplier[cardNum] = 1
		} else {
			cardWinMultiplier[cardNum]++
		}

		line := scanner.Text()

		parts := strings.Split(line, ": ")

		winnersAndNums := strings.Split(parts[1], "|")

		winnersStr := winnersAndNums[0]
		winners := strings.Split(winnersStr, " ")

		numsStr := winnersAndNums[1]
		nums := strings.Split(numsStr, " ")

		winnersMap := map[string]bool{}

		for _, winner := range winners {
			w := strings.TrimSpace(winner)
			if w == "" {
				continue
			}
			winnersMap[w] = true
		}

		winnerCount := 0

		for _, num := range nums {
			if winnersMap[strings.TrimSpace(num)] {
				// fmt.Println("winner", num)
				winnerCount++
			}
		}

		if winnerCount > 0 {
			// n wins leads to the n next cards being multiplied an extra time
			// a multiplier m for a given card means that any wins gives you m more
			// multipliers for the next n cards

			// fmt.Printf("%d instances of card %d won %d times\n", cardWinMultiplier[cardNum], cardNum, winnerCount)

			multiplier := cardWinMultiplier[cardNum]
			multiplerResult := multiplier

			for i := 0; i < winnerCount; i++ {
				cardWinMultiplier[cardNum+i+1] += multiplerResult
			}
		}

		// fmt.Printf("cardWinMultiplier: %+v\n", cardWinMultiplier)
	}

	for _, multiplier := range cardWinMultiplier {
		total += multiplier
	}

	return strconv.Itoa(total)
}
