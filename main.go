package main

import (
	"fmt"
	"github.com/aPruner/my-fridge-server/db"
	"github.com/aPruner/my-fridge-server/server"
	"github.com/joho/godotenv"
	"log"
	"net/http"
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
	fmt.Println(connString)

	gqlServer := server.Create()
	// TODO: Add server port and host to .env
	err = http.ListenAndServe("localhost:3000", gqlServer)
	if err != nil {
		log.Fatalf("There was an error starting the server")
	}
	log.Printf("Server is now listening on port 3000")

}
