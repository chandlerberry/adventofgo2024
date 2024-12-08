package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

func GetSessionCookie() ([]*http.Cookie, error) {
	sessionId, err := os.ReadFile("session.txt")
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: string(sessionId),
	}

	return []*http.Cookie{cookie}, nil
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

	jar.SetCookies(baseurl, sessionCookie)

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

	// TODO: write out input to a file, first checking if day input does not exist
	// TODO: read the file line-by-line to work with in problems

	input, err := GetDailyInput(inputUrl)
	if err != nil {
		fmt.Println("could not get daily input: ", err)
	}

	dayOnePartOne(input)
}

func dayOnePartOne(input string) {
	input_example, err := os.ReadFile("input1_example.txt")
	if err != nil {
		fmt.Println(err)
	}
	input_string := string(input_example)

	lines := strings.Split(input_string, "   ")

	fmt.Println(lines)
}
