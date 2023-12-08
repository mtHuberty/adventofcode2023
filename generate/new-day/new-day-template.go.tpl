package {{.PackageName}}

import (
	"fmt"

	"github.com/mtHuberty/adventofcode2023/util"
)

var lines{{.Part}} []string

func readFile{{.Part}}() {
	scanner := util.NewScanner("./days/{{.PackageName}}/input.txt")

	defer scanner.Close()

	for scanner.Scan() {
		line := scanner.Text()
		lines{{.Part}} = append(lines{{.Part}}, line)
	}
}

func SolvePart{{.Part}}() string {
	readFile{{.Part}}()
	return fmt.Sprintf("Counted %d input lines", len(lines{{.Part}}))
}
