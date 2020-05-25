package main

import (
	"github.com/aPruner/my-fridge-server/db"
	"github.com/aPruner/my-fridge-server/gql"
	"github.com/aPruner/my-fridge-server/server"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// TODO: Add server port and host to .env instead of defaulting to localhost:3000
	gqlServer := initServer()

	log.Fatal(http.ListenAndServe("localhost:3000", gqlServer))
}

func initServer() (gqlServer *server.Server) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbHostname := os.Getenv("DB_HOSTNAME")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connString := db.BuildConnString(dbHostname, dbPort, dbUser, dbPassword, dbName)

	database, err := db.Create(connString)
	if err != nil {
		log.Fatal(err)
	}
	gqlSchema := gql.CreateSchema(database)
	gqlServer = server.Create(&gqlSchema)
	return gqlServer
}
