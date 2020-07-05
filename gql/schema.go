package gql

import (
	"github.com/aPruner/my-fridge-server/db"
	"github.com/graphql-go/graphql"
	"log"
)

type BaseQuery struct {
	Query *graphql.Object
}

func CreateSchema(database *db.Db) graphql.Schema {
	baseQuery := NewBaseQuery(database)

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: baseQuery.Query,
		},
	)
	if err != nil {
		log.Fatalf("Error creating schema: %v", err)
	}
	return schema
}

func NewBaseQuery(database *db.Db) *BaseQuery {
	resolver := Resolver{database: database}
	baseQuery := BaseQuery{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						Type: graphql.NewList(User),
						Args: graphql.FieldConfigArgument{
							"username": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.UserResolver,
					},
					"foodItems": &graphql.Field{
						Type: graphql.NewList(FoodItem),
						Args: graphql.FieldConfigArgument{
							"householdId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.FoodItemResolver,
					},
					"householdId": &graphql.Field{
						Type: graphql.Int,
						Args: graphql.FieldConfigArgument{
							"userId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.HouseholdResolver,
					},
				},
			},
		),
	}
	return &baseQuery
}
