package database

import (
	"configs/configuration"
	"database/sql"
	"fmt"
)

func CreateConnection(config *configuration.Configuration) (*sql.DB, error) {

	// connectionString := "golang:sebastian12345@tcp(localhost:3306)/golang?parseTime=True"
	connectionString := fmt.Sprintf(
		"%s:sebastian12345@tcp(%s:3306)/golang?parseTime=True",
		config.DB.User,
		config.DB.Host)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
