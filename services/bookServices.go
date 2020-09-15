package services

import (
	"fmt"
	"log"
	"os"

	"github.com/mo-mirzania/books/database"
)

// BookServiceInterface interface
type BookServiceInterface interface {
	GetBooks() ([]database.Book, error)
}

// BookService struct
type BookService struct{}

// GetBooks func
func (bs *BookService) GetBooks() ([]database.Book, error) {
	//mysqlHandler, err := database.NewMySQLHandler(mysqlConnectionString())
	postgresqlHandler, err := database.NewPostgreSQLHandler(postgresqlConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	//books, err := mysqlHandler.GetBooks()
	books, err := postgresqlHandler.GetBooks()
	if err != nil {
		log.Println(err)
	}
	return books, err
}

func mysqlConnectionString() string {
	user := os.Getenv("SQL_USER")
	if len(user) == 0 {
		user = "root"
	}
	password := os.Getenv("SQL_PASSWORD")
	if len(password) == 0 {
		password = "mo0zakhraf"
	}
	host := os.Getenv("SQL_HOST")
	if len(host) == 0 {
		host = "localhost"
	}
	port := os.Getenv("SQL_PORT")
	if len(port) == 0 {
		port = "3306"
	}
	db := os.Getenv("SQL_DB")
	if len(db) == 0 {
		db = "books"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
}

func postgresqlConnectionString() string {
	user := os.Getenv("PG_USER")
	if len(user) == 0 {
		user = "postgres"
	}
	password := os.Getenv("PG_PASSWORD")
	if len(password) == 0 {
		password = "mo0zakhraf"
	}
	host := os.Getenv("PG_HOST")
	if len(host) == 0 {
		host = "localhost"
	}
	port := os.Getenv("PG_PORT")
	if len(port) == 0 {
		port = "5432"
	}
	db := os.Getenv("PG_DB")
	if len(db) == 0 {
		db = "books"
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db)
}
