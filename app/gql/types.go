package gql

import (
	"github.com/graphql-go/graphql"
)

var User = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var FoodItem = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FoodItem",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"category": &graphql.Field{
				Type: graphql.String,
			},
			"amount": &graphql.Field{
				Type: graphql.String,
			},
			"householdId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var Household = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Household",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"city": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ShoppingList = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ShoppingList",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"householdId": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)
