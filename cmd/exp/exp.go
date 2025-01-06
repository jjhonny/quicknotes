package main

import (
	"html/template"
	"os"
)

type TemplateData struct {
	Nome string
}

func main() {
	t, err := template.New("teste").Parse("<h1>Hello {{ .Nome }}</h1>")
	if err != nil {
		panic(err)
	}

	data := TemplateData{Nome: "Jhonny"}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
