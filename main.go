package main

import (
	"fmt"
	"github.com/aPruner/my-fridge-server/db"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHostname := os.Getenv("DB_HOSTNAME")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Error converting dbPort")
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	// Build the conn string from env vars
	connString := db.BuildConnString(dbHostname, dbPort, dbUser, dbName)
	fmt.Printf(connString)
}
