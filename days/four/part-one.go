// template.go.tpl

package four

import (
	"math"
	"strconv"
	"strings"

	"github.com/mtHuberty/adventofcode2023/util"
)

func SolvePartOne() string {
	scanner := util.NewScanner("./days/four/input.txt")

	defer scanner.Close()

	total := 0

	for scanner.Scan() {
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
				//fmt.Println("winner", num)
				winnerCount++
			}
		}

		if winnerCount > 0 {
			//fmt.Println("winnerCount", winnerCount)
			cardTotal := math.Pow(2, float64(winnerCount-1))
			//fmt.Println("cardTotal", cardTotal)

			total += int(cardTotal)
		}
	}

	return strconv.Itoa(total)
}
