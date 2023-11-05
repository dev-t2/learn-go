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

	text = "{{if .}}Action: {{.}}\n{{end}}"

	executeTemplate(text, true)
	// executeTemplate(text, false)

	fmt.Println()

	text = "{{range .}}Action: {{.}}\n{{end}}"

	executeTemplate(text, []string{"a", "b", "c"})
	// executeTemplate(text, []string{})
	// executeTemplate(text, nil)

	fmt.Println()
}