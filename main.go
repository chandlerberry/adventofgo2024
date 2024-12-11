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
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

var (
	inputYear = flag.Int("y", 2024, "Input Year")
	inputDay  = flag.Int("d", 1, "Input Day")
)

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func DayOne(input *bufio.Scanner) error {
	locationListOne := make([]int, 1000)
	locationListTwo := make([]int, 1000)

	i := 0
	for input.Scan() {
		locationInts := SliceAtoi(strings.Split(input.Text(), "   "))
		locationListOne[i] = locationInts[0]
		locationListTwo[i] = locationInts[1]

		i++
	}

	if len(locationListOne) != len(locationListTwo) {
		return errors.New("something wrong with file reading, location lists are not the same length")
	}

	slices.Sort(locationListOne)
	slices.Sort(locationListTwo)

	sum := 0
	for i = 0; i < 1000; i++ {
		x := locationListOne[i] - locationListTwo[i]
		sum += Abs(x)
	}

	fmt.Printf("%d\n", sum)
	return nil
}

func main() {
	// parse args
	flag.Parse()

	// get filepath
	pwd, pwdErr := os.Getwd()
	check(pwdErr)

	inputFilename := fmt.Sprintf("%s/inputs/day%d.txt", pwd, *inputDay)

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

	inputFile, openFileErr := os.Open(inputFilename)
	check(openFileErr)
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	day := strconv.Itoa(*inputDay)

	switch day {
	case "1":
		dayOneErr := DayOne(scanner)
		check(dayOneErr)
	case "2":
		fmt.Println("unimplemented")
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
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
