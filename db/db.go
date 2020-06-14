package db

import (
	"github.com/go-pg/pg/v10"
	"log"
)

type Db struct {
	*pg.DB
}

func BuildDbOptions(host string, port string, user string, password string, dbName string) pg.Options {
	return pg.Options{
		Addr: port,
		User: user,
		Password: password,
		Database: dbName,
	}
}

func Create(options pg.Options) (Db, error) {
	db := pg.Connect(&options)
	if err := db.Ping(nil); err != nil {
		return Db{nil}, err
	}
	return Db{db}, nil
}

func (d *Db) GetUsersByUsername(username string) []User {
	var users []User
	_, err := d.Query(pg.Scan(&users), "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		log.Printf("There were errors in the GetUsersByUsername query: %v", err)
	}
	return users
}
