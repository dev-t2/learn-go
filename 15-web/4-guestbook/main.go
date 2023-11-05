package main

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
	html, err := template.ParseFiles("15-web/4-guestbook/view.html")

	check(err)

	err = html.Execute(res, nil)

	check(err)
}

func main() {
	http.HandleFunc("/", viewHandler)

	err := http.ListenAndServe("localhost:8080", nil)

	log.Fatal(err)
}