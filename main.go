package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"
)

func DayOne() {
	fmt.Println("Day One Here")
}

func main() {
	inputYear := flag.Int("y", 2024, "Input Year")
	inputDay := flag.Int("i", 1, "Input Day")
	flag.Parse()

	// get filepath
	pwd, pwdErr := os.Getwd()
	check(pwdErr)

	inputFileName := fmt.Sprintf("%s/inputs/day%d.txt", pwd, *inputDay)

	// check if file exists, if file does not exist, download the file from https://adventofcode.com
	if _, err := os.Stat(inputFileName); err != nil {
		baseurl, urlErr := url.Parse("https://adventofcode.com")
		check(urlErr)

		jar, jarErr := cookiejar.New(nil)
		check(jarErr)

		sessionCookie, getSessionCookieErr := GetSessionCookie()
		check(getSessionCookieErr)

		jar.SetCookies(baseurl, sessionCookie)
		client := &http.Client{
			Timeout: time.Duration(30) * time.Second,
			Jar:     jar,
		}

		inputUrl := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", *inputYear, *inputDay)
		getInputErr := GetDailyInput(inputUrl, inputFileName, client)
		check(getInputErr)
	}

}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
}

func GetSessionCookie() ([]*http.Cookie, error) {
	sessionId, err := os.ReadFile("session.txt")
	check(err)

	cookie := &http.Cookie{
		Name:  "session",
		Value: string(sessionId),
	}

	return []*http.Cookie{cookie}, nil
}

func GetDailyInput(inputUrl string, inputFileName string, client *http.Client) error {
	resp, err := client.Get(inputUrl)
	check(err)

	respBytes, err := io.ReadAll(resp.Body)
	check(err)

	inputFileWriteErr := os.WriteFile(inputFileName, respBytes, 0644)
	check(inputFileWriteErr)

	return nil
}
