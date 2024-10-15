package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func InitDB() {
	var err error
	connStr := os.Getenv("DATABASE_URL")
	Conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}

	_, err = Conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS transaction(id SERIAL PRIMARY KEY, date TEXT NOT NULL, transaction REAL);")
	if err != nil {
		panic(err)
	}

	rows, err := Conn.Query(context.Background(), "SELECT * FROM transaction")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var date string
		var transaction float64
		if err := rows.Scan(&id, &date, &transaction); err != nil {
			panic(err)
		}
		fmt.Printf("%d | %s | %f\n", id, date, transaction)
	}

	fmt.Println("you are connected")
}

func CloseDB() {
	if Conn != nil {
		Conn.Close(context.Background())
	}
}
