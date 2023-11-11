package main

import (
	"errors"
	"fmt"
	"net/http"
)

func hitUrl(url string) error {
	fmt.Println(url)

	res, err := http.Get(url)

	if err != nil || res.StatusCode >= 400{
		return errors.New("Request Failed")
	}

	return nil
}

func main() {
	urls := []string{
		"https://www.google.co.kr",
		"https://www.naver.com",
		"https://www.youtube.com",
		"https://www.wavve.com",
		"https://github.com",
		"https://go.dev",
	}

	for _, url := range urls {
		hitUrl(url)
	}
}