package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

// arquive for examples
func main() {
	var err error
	dbURL := "postgres://postgres:secret@localhost:5434/postgres"
	conn, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexão com o banco foi efetuada com sucesso")
	defer conn.Close(context.Background())

	createTable()
}

func createTable() {
	query := `
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT,
			author TEXT NOT NULL
		);
	`
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table posts created")
}
