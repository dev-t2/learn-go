package main

import (
	"fmt"
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

	fmt.Println()

	text = "Template Start\nAction: {{.}}\nTemplate End\n"

	tmpl, err = template.New("test").Parse(text)

	check(err)

	err = tmpl.Execute(os.Stdout, "A")

	check(err)

	fmt.Println()

	err = tmpl.Execute(os.Stdout, 15)

	check(err)

	fmt.Println()

	err = tmpl.Execute(os.Stdout, true)

	check(err)
}