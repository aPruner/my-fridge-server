package db

import (
	"database/sql"
	"fmt"
)

type db struct {
	*sql.DB
}

func BuildConnString(host string, port int, user string, dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbName)
}
