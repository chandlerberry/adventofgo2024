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

func GetSessionCookie() (*http.Cookie, error) {
	sessionId, err := os.ReadFile("session.txt")
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:  "session",
		Value: string(sessionId),
	}

	return &cookie, nil
}

func GetDailyInput(inputUrl string) (string, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("could not create cookie jar: ", err)
	}

	baseurl, err := url.Parse("https://adventofcode.com")
	if err != nil {
		fmt.Println("could not parse url provided: ", err)
	}

	sessionCookie, err := GetSessionCookie()
	if err != nil {
		fmt.Println("could not get session cookie: ", err)
	}

	jar.SetCookies(baseurl, []*http.Cookie{sessionCookie})

	client := &http.Client{
		Timeout: time.Duration(30) * time.Second,
		Jar:     jar,
	}

	resp, err := client.Get(inputUrl)
	if err != nil {
		fmt.Println("error making request: ", err)
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body: ", err)
	}

	return string(respBytes), nil
}

func main() {
	inputYear := flag.Int("y", 2024, "Input Year")
	inputDay := flag.Int("i", 1, "Input Day")

	flag.Parse()

	inputUrl := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", *inputYear, *inputDay)

	fmt.Println(inputUrl)

	input, err := GetDailyInput(inputUrl)
	if err != nil {
		fmt.Println("could not get daily input: ", err)
	}

	fmt.Println(input)
}
