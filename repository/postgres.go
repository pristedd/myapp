package repository

import (
	"github.com/gocraft/dbr"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB() (*dbr.Connection, error) {
	db, err := dbr.Open(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=postgres password=qwerty sslmode=disable",
		nil)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
