package models

import "database/sql"

var DB *sql.DB

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func AllBooks() ([]Book, error) {
	rows, err := DB.Query(`SELECT * FROM books`)
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
