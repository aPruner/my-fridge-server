package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func BuildConnString(host string, port int, user string, password string ,dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
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
