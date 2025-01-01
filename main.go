package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	fmt.Println("Servidor rodando na porta 5000")

	//http.HandleFunc("/hello", helloHandler)
	http.Handle("/hello", http.HandlerFunc(helloHandler))

	http.ListenAndServe(":5000", nil)
}
