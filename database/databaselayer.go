package database

import "errors"

const (
	// MYSQL DB
	MYSQL int = iota
	// POSTGRESQL DB
	POSTGRESQL
	// MONGODB DB
	MONGODB
)

// Book struct
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

// BooksDBHandler interface
type BooksDBHandler interface {
	GetBooks() ([]Book, error)
}

// GetDBHandler func
func GetDBHandler(dbType int, connString string) (BooksDBHandler, error) {
	switch dbType {
	case MYSQL:
		return NewMySQLHandler(connString)
	case POSTGRESQL:
		return NewPostgreSQLHandler(connString)
		// case MONGODB:
		// 	return NewMongoDBHandler(connString)
	}
	return nil, errors.New("Invalid DB Type")
}
