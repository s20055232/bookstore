package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/s20055232/bookstore/models"
)

// Use a struct to consolidate all dependencies in one location.
type Env struct {
	// Replace the reference to models.BookModel with an interface
	// describing its methods instead. All the other code remains exactly
	// the same.
	books interface {
		All() ([]models.Book, error)
	}
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:pass@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// Set the connection pool into Env struct
	env := &Env{books: models.BookModel{DB: db}}
	// Use dependency injection here
	http.HandleFunc("/books", env.booksIndex)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	// Use Dependency injection here.
	bks, err := env.books.All()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
