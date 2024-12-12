package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"

	"github.com/chandlerberry/adventofgo2024/days"
	"github.com/chandlerberry/adventofgo2024/helpers"
)

var (
	inputYear   = flag.Int("y", 2024, "Input Year")
	inputDay    = flag.Int("d", 1, "Input Day")
	inputFolder = flag.String("f", "inputs", "Folder in project directory to get input")
)

func main() {
	flag.Parse()

	pwd, pwdErr := os.Getwd()
	helpers.Check(pwdErr)

	inputFilename := fmt.Sprintf("%s/data/%s/day%d.txt", pwd, *inputFolder, *inputDay)

	if _, err := os.Stat(inputFilename); err != nil {
		fmt.Println("Cannot find input file, downloading...")

		baseurl, urlErr := url.Parse("https://adventofcode.com")
		helpers.Check(urlErr)

		jar, jarErr := cookiejar.New(nil)
		helpers.Check(jarErr)

		sessionCookiePath := fmt.Sprintf("%s/.env/session.txt", pwd)
		sessionCookie, getSessionCookieErr := helpers.GetSessionCookies(sessionCookiePath)
		helpers.Check(getSessionCookieErr)

		jar.SetCookies(baseurl, sessionCookie)
		client := &http.Client{
			Timeout: time.Duration(30) * time.Second,
			Jar:     jar,
		}

		inputUrl := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", *inputYear, *inputDay)
		downloadInputErr := helpers.DownloadDailyInput(inputUrl, inputFilename, client)
		helpers.Check(downloadInputErr)
	}

	inputFile, openFileErr := os.Open(inputFilename)
	helpers.Check(openFileErr)
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	day := *inputDay
	switch day {
	case 1:
		days.DayOne(scanner)
	case 2:
		days.DayTwo(scanner)
	default:
		fmt.Println("Unimplemented")
	}

}
