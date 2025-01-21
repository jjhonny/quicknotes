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
	insertPost()
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
	fmt.Println("Tabela posts criada com sucesso")
}

// INSERÇÃO CONTRA SQL INJECTION
func insertPost() {
	title := "Post 1"
	content := "Contéudo do post 1"
	author := "robson"
	query := `
		INSERT INTO posts (
			title,
			content,
			author	
		)
		values ($1, $2, $3)
	`
	_, err := conn.Exec(context.Background(), query, title, content, author)
	if err != nil {
		panic(err)
	}
	fmt.Println("Post criado com sucesso")
}
