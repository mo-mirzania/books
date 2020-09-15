package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// PostgreSQLHandler struct
type PostgreSQLHandler struct {
	*SQLHandler
}

// NewPostgreSQLHandler func
func NewPostgreSQLHandler(connString string) (*PostgreSQLHandler, error) {
	db, err := sql.Open("postgres", connString)
	return &PostgreSQLHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
