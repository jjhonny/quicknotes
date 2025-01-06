package main

import (
	"html/template"
	"os"
)

type TemplateData struct {
	Nome  string
	Idade int
}

func main() {
	t, err := template.ParseFiles("./cmd/hello.html")
	if err != nil {
		panic(err)
	}

	data := TemplateData{Nome: "Jhonny", Idade: 22}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
