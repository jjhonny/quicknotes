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
	t, err := template.ParseFiles("layout1.html", "home.html", "footer.html", "header.html")

	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "layout1.html", nil)
	if err != nil {
		panic(err)
	}
}
