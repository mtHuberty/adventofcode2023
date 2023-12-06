package three

import (
	"strconv"
)

func SolvePartTwo() string {
	total := 0

	for y, row := range sm.matrix {
		for x, char := range row {
			if char == "*" {
				//fmt.Printf("special char: %s found at: %d, %d\n", char, x, y)

				coords := getAdjacentCoords(x, y)

				alreadySeenCoords := map[[2]int]bool{}

				numFoundCount := 0

				nums := []int{}

				for _, coord := range coords {
					if _, ok := alreadySeenCoords[coord]; ok {
						continue
					}

					alreadySeenCoords[coord] = true

					if getCharAtCoords(coord[0], coord[1]) == "." {
						continue
					}

					num, foundAt := getFullNumAtCoord(coord[0], coord[1])
					nums = append(nums, num)
					numFoundCount++

					if len(foundAt) > 0 {
						for _, x := range foundAt {
							alreadySeenCoords[[2]int{x, coord[1]}] = true
						}
					}
				}

				if numFoundCount == 2 {
					total += nums[0] * nums[1]
				}
			}
		}
	}

	return strconv.Itoa(total)
}
