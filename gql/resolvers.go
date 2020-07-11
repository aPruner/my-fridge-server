package gql

import (
	"fmt"
	"github.com/aPruner/my-fridge-server/db"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	database *db.Db
}

func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	// Type-check the name
	username, ok := p.Args["username"].(string)
	if ok {
		users := r.database.GetUsersByUsername(username)
		return users, nil
	}
	err := fmt.Errorf("type-checking error: username was not a string")
	return nil, err
}

func (r *Resolver) FoodItemResolver(p graphql.ResolveParams) (interface{}, error) {
	// Type-check the householdId
	householdId, ok := p.Args["householdId"].(int)
	if ok {
		foodItems := r.database.GetFoodItemsByHousehold(householdId)
		return foodItems, nil
	}
	err := fmt.Errorf("type-checking error: householdId was not an int")
	return nil, err
}

func (r *Resolver) HouseholdResolver(p graphql.ResolveParams)  (interface{}, error) {
	// Type-check the userId
	userId, ok := p.Args["userId"].(int)
	if ok {
		householdId := r.database.GetHouseholdIdByUserId(userId)
		return householdId, nil
	}
	err := fmt.Errorf("type-checking error: userId was not an int")
	return nil, err
}
