package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	Count    int
	Contents []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(filename string) []string {
	var lines []string

	file, err := os.Open(filename)

	if os.IsNotExist(err) {
		return nil
	}

	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	check(scanner.Err())

	return lines
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
	signatures := getStrings("15-web/7-guestbook/signatures.txt")

	html, err := template.ParseFiles("15-web/7-guestbook/view.html")

	check(err)

	guestbook := Guestbook{Count: len(signatures), Contents: signatures}

	err = html.Execute(res, guestbook)

	check(err)
}

func newHandler(res http.ResponseWriter, req *http.Request) {
	html, err := template.ParseFiles("15-web/7-guestbook/new.html")

	check(err)

	err = html.Execute(res, nil)

	check(err)
}

func createHandler(res http.ResponseWriter, req *http.Request) {
	content := req.FormValue("content")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE

	file, err := os.OpenFile("15-web/7-guestbook/signatures.txt", options, os.FileMode(0600))

	check(err)

	_, err = fmt.Fprintln(file, content)

	check(err)

	err = file.Close()

	check(err)

	http.Redirect(res, req, "/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", viewHandler)

	http.HandleFunc("/new", newHandler)

	http.HandleFunc("/create", createHandler)

	err := http.ListenAndServe("localhost:8080", nil)

	log.Fatal(err)
}