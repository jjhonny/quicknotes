package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Handler")
}

func olaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ola Handler")
}

func main() {
	fmt.Println("Servidor rodando na porta 5000")
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/ola/", olaHandler)
	mux.HandleFunc("/ola/mundo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ola Mundo Handler")
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Handler")
	})

	mux.HandleFunc("ola.pessoas.com/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ola Pessoas Handler")
	})

	http.ListenAndServe(":5000", mux)
}
