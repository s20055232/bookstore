package models

import (
	"database/sql"
	"log"
)

// Turn it into lowercase bacause we don't need to export it.
var db *sql.DB

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:pass@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db.Ping()
}

func AllBooks() ([]Book, error) {
	rows, err := db.Query(`SELECT * FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bks []Book

	for rows.Next() {
		var bk Book

		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}
