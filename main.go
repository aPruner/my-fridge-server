package main

import (
	"fmt"
	"github.com/aPruner/my-fridge-server/app/db"
	"github.com/aPruner/my-fridge-server/app/gql"
	"github.com/aPruner/my-fridge-server/app/server"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	gqlServer := initServer()
	serverEnv := os.Getenv("SERVER_ENV")
	serverHost := os.Getenv(fmt.Sprintf("SERVER_HOST_%s", serverEnv))
	log.Printf("Server created, now listening at localhost:8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:8080", serverHost), gqlServer))
}

func initServer() (gqlServer *server.Server) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env variables: %v", err)
	}

	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOSTNAME")

	connString := db.BuildDbOptions(dbHost, dbPort, dbUser, dbPassword, dbName)

	database, err := db.Create(connString)
	if err != nil {
		log.Fatalf("Error creating the database: %v", err)
	}
	gqlSchema := gql.CreateSchema(database)
	gqlServer = server.Create(&gqlSchema)
	return gqlServer
}
