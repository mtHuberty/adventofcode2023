package five

import (
	"fmt"

	"github.com/mtHuberty/adventofcode2023/util"
)

var linesTwo []string

func init() {
	scanner := util.NewScanner("./days/five/input.txt")

	defer scanner.Close()

	for scanner.Scan() {
		line := scanner.Text()
		linesTwo = append(linesTwo, line)
	}
}

func SolvePartTwo() string {
	return fmt.Sprintf("Counted %d input lines", len(linesTwo))
}
