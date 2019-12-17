package main

import (
	"database/sql"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// SQLHandler for access db
var SQLHandler *sql.DB

func sqlInit() {
	conn, err := sql.Open("mysql", "docker:docker@/tweeet")
	if err != nil {
		panic(err.Error)
	}
	fmt.Println("successflly get db connect")
	SQLHandler = conn
}

func getMessage() {

	// SELECT
	rows, err := SQLHandler.Query("SELECT * FROM messages")

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.Name, &m.Message); err != nil {
			fmt.Println(err)
		}
		// names = append(names, name)
		fmt.Println("Log MessgeList", m)
	}
}

func main() {
	sqlInit()
	getMessage()
}

// Message struct
type Message struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

