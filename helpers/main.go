package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

func Check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
}

func SliceAtoi(stringSlice []string) []int {
	ints := make([]int, len(stringSlice))

	for i, s := range stringSlice {
		ints[i], _ = strconv.Atoi(s)
	}

	return ints
}

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}

	return x
}

func GetSessionCookies(sessionCookiePath string) ([]*http.Cookie, error) {
	sessionId, err := os.ReadFile(sessionCookiePath)
	Check(err)

	cookie := &http.Cookie{
		Name:  "session",
		Value: string(sessionId),
	}

	return []*http.Cookie{cookie}, nil
}

func DownloadDailyInput(inputUrl string, inputFileName string, client *http.Client) error {
	resp, err := client.Get(inputUrl)
	Check(err)

	respBytes, err := io.ReadAll(resp.Body)
	Check(err)

	inputFileWriteErr := os.WriteFile(inputFileName, respBytes, 0644)
	Check(inputFileWriteErr)

	return nil
}
