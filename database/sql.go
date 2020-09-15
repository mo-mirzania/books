package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// SQLHandler struct
type SQLHandler struct {
	*sql.DB
}

// GetBooks func
func (handler *SQLHandler) GetBooks() ([]Book, error) {
	rows, err := handler.Query(fmt.Sprintf("SELECT * FROM books;"))
	if err != nil {
		return nil, err
	}
	var books []Book
	var book Book
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			log.Println(err)
			continue
		}
		books = append(books, book)
	}
	return books, err
}
