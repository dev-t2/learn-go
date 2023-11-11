package main

import (
	"errors"
	"fmt"
	"net/http"
)

func hitUrl(url string) error {
	fmt.Println("Checking:", url)

	res, err := http.Get(url)

	if err != nil || res.StatusCode >= 400{
		fmt.Println(err, res.StatusCode)

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
		"https://www.inflearn.com",
		"https://nomadcoders.co",
		"https://fastcampus.co.kr",
		"https://go.dev",
	}

	results := make(map[string]string)

	for _, url := range urls {
		err := hitUrl(url)

		if err != nil {
			results[url] = "Failure"
		} else {
			results[url] = "Success"
		}
	}

	fmt.Println()

	for url, result := range results {
		fmt.Println(url, result)
	}
}