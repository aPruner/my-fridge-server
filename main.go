package main

import (
	"github.com/aPruner/my-fridge-server/db"
	"github.com/aPruner/my-fridge-server/gql"
)

func main() {
	gql.GqlTest()
	db.PrintDB()
}
