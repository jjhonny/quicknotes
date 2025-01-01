package main

import (
	"fmt"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listagem de notas...")
}

func noteView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Exibindo uma nota específica...")
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Criando uma nova nota...")
}

func main() {
	fmt.Println("Servidor rodando na porta 5000")
	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":5000", mux)
}
