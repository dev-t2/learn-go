package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url string
	status string
}

func hitUrl(channel chan<- result, url string) {
	res, err := http.Get(url)
	status := "Success"

	if err != nil || res.StatusCode >= 400{
		status = "Failure"
	} 

	channel<-result{url: url, status: status}
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

	channel := make(chan result)

	for _, url := range urls {
		go hitUrl(channel, url)
	}

	results := make(map[string] string)

	for i := 0; i < len(urls); i++ {
		result := <-channel

		results[result.url] = result.status
	}

	fmt.Println()

	for url, status := range results {
		fmt.Println(url, status)
	}
}