package main

import (
	"fmt"
	"html/template"
	"os"
)

type TemplateData struct {
	Nome  string
	Idade int
}

func main() {
	t, err := template.ParseFiles("b.html", "a.html")

	fmt.Println(t.Name())

	fmt.Println(t.DefinedTemplates())

	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "a.html", nil)
	if err != nil {
		panic(err)
	}
}
