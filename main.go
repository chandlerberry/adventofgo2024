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

func main() {
	inputYear := flag.Int("y", 2024, "Input Year")
	inputDay := flag.Int("i", 1, "Input Day")
	flag.Parse()

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

	getInputErr := GetDailyInput(inputYear, inputDay, client)
	check(getInputErr)
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

func GetDailyInput(inputYear *int, inputDay *int, client *http.Client) error {
	inputUrl := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", *inputYear, *inputDay)

	resp, err := client.Get(inputUrl)
	check(err)

	respBytes, err := io.ReadAll(resp.Body)
	check(err)

	pwd, pwdErr := os.Getwd()
	check(pwdErr)

	inputFileName := fmt.Sprintf("%s/inputs/day%d.txt", pwd, *inputDay)

	inputFileWriteErr := os.WriteFile(inputFileName, respBytes, 0644)
	check(inputFileWriteErr)

	return nil
}
