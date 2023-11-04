package main

import (
	"log"
	"net/http"
)

func write(res http.ResponseWriter, message string) {
	_, err := res.Write([]byte(message))

	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(res http.ResponseWriter, req *http.Request) {
	write(res, "Hello, Web!")
}

func frenchHandler(res http.ResponseWriter, req *http.Request) {
	write(res, "Salut, Web!")
}

func hindiHandler(res http.ResponseWriter, req *http.Request) {
	write(res, "Namaste, Web!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)

	http.HandleFunc("/salut", frenchHandler)

	http.HandleFunc("/namaste", hindiHandler)

	err := http.ListenAndServe("localhost:8080", nil)

	log.Fatal(err)
}