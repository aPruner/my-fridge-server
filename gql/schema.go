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
		log.Fatal(err)
	}
	return schema
}

func NewBaseQuery(database *db.Db) *BaseQuery {
	resolver := Resolver{database: database}
	baseQuery := BaseQuery{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Fields: graphql.Fields{
					"users": &graphql.Field{
						Type: graphql.NewList(User),
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.UserResolver,
					},
				},
			},
		),
	}
	return &baseQuery
}
