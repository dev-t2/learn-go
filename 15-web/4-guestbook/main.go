package main

import (
	"bufio"
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
	signatures := getStrings("15-web/4-guestbook/signatures.txt")

	html, err := template.ParseFiles("15-web/4-guestbook/view.html")

	check(err)

	guestbook := Guestbook{Count: len(signatures), Contents: signatures}

	err = html.Execute(res, guestbook)

	check(err)
}

func main() {
	http.HandleFunc("/", viewHandler)

	err := http.ListenAndServe("localhost:8080", nil)

	log.Fatal(err)
}