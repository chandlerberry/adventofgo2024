package days

import (
	"bufio"
	"fmt"
	"sort"
	"strings"

	"github.com/chandlerberry/adventofgo2024/helpers"
)

func DayTwo(input *bufio.Scanner) {
	partOne := 0
	partTwo := 0

	for input.Scan() {
		levels := helpers.SliceAtoi(strings.Split(input.Text(), " "))

		if isSafe(levels) {
			partOne++
		}

		if isSafeWithDampener(levels) {
			partTwo++
		}
	}

	fmt.Printf("Part One: %d\n", partOne)
	fmt.Printf("Part Two: %d\n", partTwo)
}

func isSafe(levels []int) bool {
	isSortedAsc := sort.SliceIsSorted(levels, func(i int, j int) bool {
		return levels[i] < levels[j]
	})

	isSortedDesc := sort.SliceIsSorted(levels, func(i int, j int) bool {
		return levels[j] < levels[i]
	})

	if !isSortedAsc && !isSortedDesc {
		return false
	}

	for i, v := range levels {
		if i == len(levels)-1 {
			break
		}

		levelChange := helpers.Abs(v - levels[i+1])

		if levelChange <= 0 || levelChange > 3 {
			return false
		}
	}

	return true
}

func isSafeWithDampener(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	if !isSafe(levels) {
		for i := range levels {
			dampenedList := append([]int(nil), levels...)
			dampenedList = append(dampenedList[:i], dampenedList[i+1:]...)

			if isSafe(dampenedList) {
				return true
			}
		}
	}

	return false
}
