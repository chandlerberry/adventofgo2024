package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

var (
	inputYear   = flag.Int("y", 2024, "Input Year")
	inputDay    = flag.Int("d", 1, "Input Day")
	inputFolder = flag.String("f", "inputs", "Folder in project directory to get input")
)

func DayTwo(input *bufio.Scanner) error {
	partOne := 0
	partTwo := 0

	for input.Scan() {
		levels := SliceAtoi(strings.Split(input.Text(), " "))
		safety := 1

		for i, v := range levels {
			if i == len(levels)-1 {
				break
			}

			levelChange := Abs(v - levels[i+1])

			if levelChange <= 0 || levelChange > 3 {
				safety = 0
				break
			}
		}

		if safety == 0 {
			continue
		}

		isSortedAsc := sort.SliceIsSorted(levels, func(i int, j int) bool {
			return levels[i] < levels[j]
		})

		isSortedDesc := sort.SliceIsSorted(levels, func(i int, j int) bool {
			return levels[j] < levels[i]
		})

		if !isSortedAsc && !isSortedDesc {
			continue
		}

		partOne++
	}

	fmt.Printf("Part One: %d\n", partOne)
	fmt.Printf("Part Two: %d\n", partTwo)
	return nil
}

func DayOne(input *bufio.Scanner) error {
	locationListOne := make([]int, 1000)
	locationListTwo := make([]int, 1000)

	line := 0
	for input.Scan() {
		locationInts := SliceAtoi(strings.Split(input.Text(), "   "))
		locationListOne[line] = locationInts[0]
		locationListTwo[line] = locationInts[1]

		line++
	}

	if len(locationListOne) != len(locationListTwo) {
		return errors.New("something wrong with file reading, location lists are not the same length")
	}

	slices.Sort(locationListOne)
	slices.Sort(locationListTwo)

	partOne := 0
	partTwo := 0
	for i := 0; i < 1000; i++ {
		x := locationListOne[i] - locationListTwo[i]
		partOne += Abs(x)

		frequency := 0
		for j := 0; j < 1000; j++ {
			if locationListOne[i] == locationListTwo[j] {
				frequency++
			}
		}
		score := frequency * locationListOne[i]
		partTwo += score

	}

	fmt.Printf("%d\n", partOne)
	fmt.Printf("%d\n", partTwo)

	return nil
}

func main() {
	// parse args
	flag.Parse()

	// get filepath
	pwd, pwdErr := os.Getwd()
	check(pwdErr)

	inputFilename := fmt.Sprintf("%s/%s/day%d.txt", pwd, *inputFolder, *inputDay)

	// check if file exists, if file does not exist, download the file from https://adventofcode.com
	if _, err := os.Stat(inputFilename); err != nil {
		fmt.Println("Cannot find input file, downloading...")

		baseurl, urlErr := url.Parse("https://adventofcode.com")
		check(urlErr)

		jar, jarErr := cookiejar.New(nil)
		check(jarErr)

		sessionCookie, getSessionCookieErr := GetSessionCookies()
		check(getSessionCookieErr)

		jar.SetCookies(baseurl, sessionCookie)
		client := &http.Client{
			Timeout: time.Duration(30) * time.Second,
			Jar:     jar,
		}

		inputUrl := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", *inputYear, *inputDay)
		downloadInputErr := DownloadDailyInput(inputUrl, inputFilename, client)
		check(downloadInputErr)
	}

	// open the input file
	inputFile, openFileErr := os.Open(inputFilename)
	check(openFileErr)
	defer inputFile.Close()

	// create the scanner over the file
	scanner := bufio.NewScanner(inputFile)

	// run the
	day := *inputDay
	switch day {
	case 1:
		dayOneErr := DayOne(scanner)
		check(dayOneErr)
	case 2:
		dayTwoErr := DayTwo(scanner)
		check(dayTwoErr)
	default:
		fmt.Println("Unimplemented")
	}

}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
}

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func GetSessionCookies() ([]*http.Cookie, error) {
	sessionId, err := os.ReadFile("session.txt")
	check(err)

	cookie := &http.Cookie{
		Name:  "session",
		Value: string(sessionId),
	}

	return []*http.Cookie{cookie}, nil
}

func DownloadDailyInput(inputUrl string, inputFileName string, client *http.Client) error {
	resp, err := client.Get(inputUrl)
	check(err)

	respBytes, err := io.ReadAll(resp.Body)
	check(err)

	inputFileWriteErr := os.WriteFile(inputFileName, respBytes, 0644)
	check(inputFileWriteErr)

	return nil
}

func SliceAtoi(stringSlice []string) []int {
	ints := make([]int, len(stringSlice))
	for i, s := range stringSlice {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}
