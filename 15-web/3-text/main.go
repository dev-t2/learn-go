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
	text := "Template Start\nTemplate End\n"
	tmpl, err := template.New("test").Parse(text)

	check(err)

	err = tmpl.Execute(os.Stdout, nil)

	check(err)

	fmt.Println()

	text = "Template Start\nAction: {{.}}\nTemplate End\n"

	tmpl, err = template.New("test").Parse(text)

	check(err)

	err = tmpl.Execute(os.Stdout, 15)

	check(err)

	fmt.Println()

	executeTemplate(text, 15)

	fmt.Println()

	text = "Template Start\n{{if .}}Action: true\n{{end}}Template End\n"

	executeTemplate(text, true)

	fmt.Println()
	
	executeTemplate(text, false)
}