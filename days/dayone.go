package days

import (
	"bufio"
	"fmt"
	"slices"
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
