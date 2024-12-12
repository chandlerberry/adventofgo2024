package days

import (
	"bufio"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/chandlerberry/adventofgo2024/helpers"
)

func DayOne(input *bufio.Scanner) {
	locationListOne, locationListTwo := make([]int, 1000), make([]int, 1000)

	line := 0
	for input.Scan() {
		locationInts := helpers.SliceAtoi(strings.Split(input.Text(), "   "))

		locationListOne[line] = locationInts[0]
		locationListTwo[line] = locationInts[1]

		line++
	}

	slices.Sort(locationListOne)
	slices.Sort(locationListTwo)

	partOne, partTwo := 0, 0

	for i := 0; i < 1000; i++ {
		x := locationListOne[i] - locationListTwo[i]
		partOne += helpers.Abs(x)

		frequency := 0
		for j := 0; j < 1000; j++ {
			if locationListOne[i] == locationListTwo[j] {
				frequency++
			}
		}

		score := frequency * locationListOne[i]
		partTwo += score

	}

	fmt.Printf("Part One: %d\n", partOne)
	fmt.Printf("Part Two: %d\n", partTwo)
}

func DayTwo(input *bufio.Scanner) {
	partOne := 0
	partTwo := 0

	for input.Scan() {
		levels := helpers.SliceAtoi(strings.Split(input.Text(), " "))
		safety := 1

		for i, v := range levels {
			if i == len(levels)-1 {
				break
			}

			levelChange := helpers.Abs(v - levels[i+1])

			if levelChange <= 0 || levelChange > 3 {
				safety = 0
				break
			}
		}

		if safety == 0 {
			continue
		}

		isSortedAsc, isSortedDesc := sort.SliceIsSorted(levels, func(i int, j int) bool {
			return levels[i] < levels[j]
		}), sort.SliceIsSorted(levels, func(i int, j int) bool {
			return levels[j] < levels[i]
		})

		if !isSortedAsc && !isSortedDesc {
			continue
		}

		partOne++
	}

	fmt.Printf("Part One: %d\n", partOne)
	fmt.Printf("Part Two: %d\n", partTwo)
}
