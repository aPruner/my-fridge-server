package db

import (
	"database/sql"
	"fmt"
)

type Db struct {
	*sql.DB
}

func BuildConnString(host string, port int, user string, dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbName)
}

func Create(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Db{db}, nil
}
