package main

import (
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
	placeholder := []byte("Signature list goes here")

	_, err := res.Write(placeholder)

	check(err)
}

func main() {
	http.HandleFunc("/", viewHandler)

	err := http.ListenAndServe("localhost:8080", nil)

	log.Fatal(err)
}