package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Db struct {
	*sql.DB
}

func BuildConnString(host string, port int, user string, password string, dbName string) string {
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

func (d *Db) GetUsersByUsername(username string) []User {
	sqlStatement, err := d.Prepare("SELECT * FROM users WHERE username=$1")
	if err != nil {
		log.Printf("GetUsersByUsername sql prep error: %v", err)
	}

	rows, err := sqlStatement.Query(username)
	if err != nil {
		log.Printf("GetUsersByUsername sql query error: %v", err)
	}

	var r User
	var users []User
	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Username,
		)
		if err != nil {
			log.Printf("GetUsersByUsername scanning rows error: %v", err)
		}
		users = append(users, r)
	}
	return users
}
