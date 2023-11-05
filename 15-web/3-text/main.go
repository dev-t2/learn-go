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

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)

	check(err)

	err = tmpl.Execute(os.Stdout, data)

	check(err)
}

func main() {
	text := "Template\n"
	tmpl, err := template.New("test").Parse(text)

	check(err)

	err = tmpl.Execute(os.Stdout, nil)

	check(err)

	fmt.Println()

	text = "Action: {{.}}\n"

	tmpl, err = template.New("test").Parse(text)

	check(err)

	err = tmpl.Execute(os.Stdout, 15)

	check(err)

	fmt.Println()

	executeTemplate(text, 15)

	fmt.Println()

	text = "{{if .}}Action: {{.}}{{end}}\n"

	executeTemplate(text, true)

	fmt.Println()

	text = "Slice: {{.}}\n{{range .}}Action: {{.}}\n{{end}}\n"

	executeTemplate(text, []string{"a", "b", "c"})
}