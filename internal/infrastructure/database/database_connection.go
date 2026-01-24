package database

import (
	"database/sql"
	"fmt"
)

func NewDatabaseConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected database successfully")
	return db, nil
}
