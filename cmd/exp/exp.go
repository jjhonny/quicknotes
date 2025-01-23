package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Post struct {
	id      int
	title   string
	content string
	author  string
}

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

	// createTable()
	// insertPost()
	// insertPostWithReturn()
	// selectById()
	selectAllPosts()
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

func insertPostWithReturn() {
	title := "Post 3"
	content := "Contéudo do post 3"
	author := "robson"
	query := `
        INSERT INTO posts (title, content, author)
        values ($1, $2, $3) RETURNING id;
    `

	row := conn.QueryRow(context.Background(), query, title, content, author)
	var id int
	if err := row.Scan(&id); err != nil {
		panic(err)
	}
	fmt.Println("Post Criado. id =", id)
}

func selectById() {
	id := 3
	var title, content, author string
	query := "SELECT title, content, author FROM posts WHERE id = $1"
	row := conn.QueryRow(context.Background(), query, id)
	err := row.Scan(&title, &content, &author)
	if err == pgx.ErrNoRows {
		fmt.Println("No post found for id = ", id)
		return
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf("Post: title=%s, content=%s, author=%s \n", title, content, author)
}

func selectAllPosts() {
	query := `
        SELECT * FROM posts
    `
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Err() != nil {
		panic(rows.Err())
	}

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.id, &post.title, &post.content, &post.author)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	for _, post := range posts {
		fmt.Println(post)
	}
}
