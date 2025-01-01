package main

import (
	"fmt"
	"net/http"
)

type WorldHandler struct{}

func (WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("World!!"))
}

type HelloHandler struct{}

func (HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	fmt.Println("Servidor rodando na porta 5000")

	hello := HelloHandler{}
	world := WorldHandler{}

	http.Handle("/hello", hello)
	http.Handle("/world", world)

	http.ListenAndServe(":5000", nil)

}
