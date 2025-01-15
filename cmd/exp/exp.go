package main

import (
	"flag"
	"fmt"
)

func main() {
	var port string
	var verbose bool
	var valor int
	flag.StringVar(&port, "port", "7000", "Porta do servidor")
	flag.BoolVar(&verbose, "v", false, "Modo verboso")
	flag.IntVar(&valor, "valor", 0, "Valor a ser exibido")

	flag.Parse()

	if verbose {
		fmt.Println("Server rodando na porta", port)
		fmt.Println(valor)
	} else {
		fmt.Println(port)
	}
}
