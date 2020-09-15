package database

import "database/sql"

type MySQLHandler struct {
	*SQLHandler
}

func NewMySQLHandler(connString string) (*MySQLHandler, error) {
	db, err := sql.Open("mysql", connString)
	return &MySQLHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
