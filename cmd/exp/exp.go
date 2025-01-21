package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// arquive for examples
func main() {
	dbURL := "postgres://postgres:secret@localhost:5434/postgres"
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexão com o banco foi efetuada com sucesso")
	defer conn.Close(context.Background())
}
