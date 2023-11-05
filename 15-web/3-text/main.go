package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	text := "Text Template\n"
	tmpl, err := template.New("test").Parse(text)

	check(err)

	err = tmpl.Execute(os.Stdout, nil)

	check(err)
}