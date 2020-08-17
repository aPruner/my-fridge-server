package gql

import (
	"github.com/aPruner/my-fridge-server/app/db"
	"github.com/graphql-go/graphql"
	"log"
)

// May not need these types but for now they are here for clarity
type BaseQuery struct {
	Query *graphql.Object
}

type BaseMutation struct {
	Mutation *graphql.Object
}

func CreateSchema(database *db.Db) graphql.Schema {
	baseQuery := CreateBaseQuery(database)
	baseMutation := CreateBaseMutation(database)

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    baseQuery.Query,
			Mutation: baseMutation.Mutation,
		},
	)
	if err != nil {
		log.Fatalf("Error creating schema: %v", err)
	}
	return schema
}

func CreateBaseQuery(database *db.Db) *BaseQuery {
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
						Resolve: resolver.UserQueryResolver,
					},
					"foodItems": &graphql.Field{
						Type: graphql.NewList(FoodItem),
						Args: graphql.FieldConfigArgument{
							"householdId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.FoodItemQueryResolver,
					},
					"householdIds": &graphql.Field{
						Type: graphql.Int,
						Args: graphql.FieldConfigArgument{
							"userId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.HouseholdIdQueryResolver,
					},
					"households": &graphql.Field{
						Type: graphql.NewList(Household),
						Args: graphql.FieldConfigArgument{
							"householdId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.HouseholdQueryResolver,
					},
					"shoppingLists": &graphql.Field{
						Type: graphql.NewList(ShoppingList),
						Args: graphql.FieldConfigArgument{
							"householdId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.ShoppingListQueryResolver,
					},
				},
			},
		),
	}
	return &baseQuery
}

func CreateBaseMutation(database *db.Db) *BaseMutation {
	resolver := Resolver{database}
	baseMutation := BaseMutation{
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Mutation",
				Fields: graphql.Fields{
					"createFoodItem": &graphql.Field{
						Type: FoodItem,
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"category": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"amount": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"unit": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"householdId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"shoppingListId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.CreateFoodItemMutationResolver,
					},
					"updateFoodItem": &graphql.Field{
						Type: FoodItem,
						// TODO: Figure out how to do optional arguments for the GQL mutations
						// TODO: Ideally, the caller should be able to update whichever fields they want, not be
						// TODO: forced to update all of them at once
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"category": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"amount": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"unit": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"householdId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"shoppingListId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.UpdateFoodItemMutationResolver,
					},
					"deleteFoodItem": &graphql.Field{
						Type: FoodItem,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.DeleteFoodItemMutationResolver,
					},
					"createShoppingList": &graphql.Field{
						Type: ShoppingList,
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"description": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"userId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"householdId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"createdAt": &graphql.ArgumentConfig{
								Type: graphql.DateTime,
							},
						},
						Resolve: resolver.CreateShoppingListResolver,
					},
					"updateShoppingList": &graphql.Field{
						Type: ShoppingList,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"description": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"userId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"householdId": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"createdAt": &graphql.ArgumentConfig{
								Type: graphql.DateTime,
							},
						},
						Resolve: resolver.UpdateShoppingListResolver,
					},
					"deleteShoppingList": &graphql.Field{
						Type: ShoppingList,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.DeleteShoppingListResolver,
					},
				},
			},
		),
	}
	return &baseMutation
}
