// template.go.tpl

package three

import (
	// "fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mtHuberty/adventofcode2023/util"
)

type schematicMatrix struct {
	matrix [][]string
}

func init() {
	scanner := util.NewScanner("./days/three/input.txt")

	defer scanner.Close()

	for scanner.Scan() {
		line := scanner.Text()

		chars := strings.Split(line, "")

		// Prepend chars so we build matrix from the bottom up
		// It's easier to reason about as a regular x,y coordinate plane
		// when +x is -> and +y is ^
		sm.matrix = append([][]string{chars}, sm.matrix...)
	}
}

func (sm schematicMatrix) rowLen() int {
	if len(sm.matrix) == 0 {
		return 0
	}

	return len(sm.matrix[0])
}

var sm schematicMatrix

// 1 3 4 . . 2 . . 3
// . . . * 2 3 4 . .
// 5 . 4 4 2 4 . . .
func SolvePartOne() string {
	scanner := util.NewScanner("./days/three/input.txt")

	defer scanner.Close()

	for scanner.Scan() {
		line := scanner.Text()

		chars := strings.Split(line, "")

		// Prepend chars so we build matrix from the bottom up
		// It's easier to reason about as a regular x,y coordinate plane
		// when +x is -> and +y is ^
		sm.matrix = append([][]string{chars}, sm.matrix...)
	}

	total := 0

	for y, row := range sm.matrix {
		for x, char := range row {
			if !strings.Contains(".1234567890", char) {
				//fmt.Printf("special char: %s found at: %d, %d\n", char, x, y)

				coords := getAdjacentCoords(x, y)

				alreadySeenCoords := map[[2]int]bool{}

				for _, coord := range coords {
					if _, ok := alreadySeenCoords[coord]; ok {
						continue
					}

					alreadySeenCoords[coord] = true

					if getCharAtCoords(coord[0], coord[1]) == "." {
						continue
					}

					num, foundAt := getFullNumAtCoord(coord[0], coord[1])

					total += num

					if len(foundAt) > 0 {
						for _, x := range foundAt {
							alreadySeenCoords[[2]int{x, coord[1]}] = true
						}
					}
				}
			}
		}
	}

	return strconv.Itoa(total)
}

func getAdjacentCoords(x, y int) [][2]int {
	coords := [][2]int{}

	// Down
	if y-1 >= 0 {
		coords = append(coords, [2]int{x, y - 1})
	}

	// Up
	if y+1 < len(sm.matrix) {
		coords = append(coords, [2]int{x, y + 1})
	}

	// Left
	if x-1 >= 0 {
		coords = append(coords, [2]int{x - 1, y})
	}

	// Right
	if x+1 < sm.rowLen() {
		coords = append(coords, [2]int{x + 1, y})
	}

	// Down-Left
	if y-1 >= 0 && x-1 >= 0 {
		coords = append(coords, [2]int{x - 1, y - 1})
	}

	// Down-Right
	if y-1 >= 0 && x+1 < sm.rowLen() {
		coords = append(coords, [2]int{x + 1, y - 1})
	}

	// Up-Left
	if y+1 < len(sm.matrix) && x-1 >= 0 {
		coords = append(coords, [2]int{x - 1, y + 1})
	}

	// Up-Right
	if y+1 < len(sm.matrix) && x+1 < sm.rowLen() {
		coords = append(coords, [2]int{x + 1, y + 1})
	}

	return coords
}

func getCharAtCoords(x, y int) string {
	return sm.matrix[y][x]
}

// Returns the number found at the coordinate, along with
// any x values that the number covers up on the plane
func getFullNumAtCoord(x, y int) (int, []int) {
	char, err := strconv.Atoi(getCharAtCoords(x, y))
	if err != nil {
		panic(err)
	}

	numStr := strconv.Itoa(char)
	foundAt := []int{}

	for i := x + 1; i < sm.rowLen(); i++ {
		char := getCharAtCoords(i, y)
		//fmt.Println("right char:", char)
		if _, err := strconv.Atoi(char); err == nil {
			foundAt = append(foundAt, i)
			numStr = numStr + char
		} else {
			break
		}
	}

	for i := x - 1; i >= 0; i-- {
		char := getCharAtCoords(i, y)
		//fmt.Println("left char:", char)
		if _, err := strconv.Atoi(char); err == nil {
			foundAt = append(foundAt, i)
			numStr = char + numStr
		} else {
			break
		}
	}

	//fmt.Println("numStr:", numStr)

	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal("error converting numStr to int:", err)
	}

	return num, foundAt
}
