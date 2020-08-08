package gql

import (
	"fmt"
	"github.com/aPruner/my-fridge-server/app/db"
	"github.com/graphql-go/graphql"
	"log"
)

type Resolver struct {
	database *db.Db
}

func (r *Resolver) UserQueryResolver(p graphql.ResolveParams) (interface{}, error) {
	// Type-check the name
	username, ok := p.Args["username"].(string)
	if ok {
		users, err := r.database.GetUsersByUsername(username)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
	err := fmt.Errorf("type-checking error: username was not a string")
	log.Print(err)
	return nil, err
}

func (r *Resolver) FoodItemQueryResolver(p graphql.ResolveParams) (interface{}, error) {
	// Type-check the householdId
	householdId, ok := p.Args["householdId"].(int)
	if ok {
		foodItems, err := r.database.GetFoodItemsByHouseholdId(householdId)
		if err != nil {
			return nil, err
		}
		return foodItems, nil
	}
	err := fmt.Errorf("type-checking error: householdId was not an int")
	log.Print(err)
	return nil, err
}

func (r *Resolver) HouseholdQueryResolver(p graphql.ResolveParams) (interface{}, error) {
	// Type-check the userId
	userId, ok := p.Args["userId"].(int)
	if ok {
		householdId, err := r.database.GetHouseholdIdByUserId(userId)
		if err != nil {
			return nil, err
		}
		return householdId, nil
	}
	err := fmt.Errorf("type-checking error: userId was not an int")
	log.Print(err)
	return nil, err
}

func (r *Resolver) CreateFoodItemMutationResolver(p graphql.ResolveParams) (interface{}, error) {
	name, nameOk := p.Args["name"].(string)
	category, categoryOk := p.Args["category"].(string)
	amount, amountOk := p.Args["amount"].(int)
	householdId, householdIdOk := p.Args["householdId"].(int)
	if nameOk && categoryOk && amountOk && householdIdOk {
		newFoodItemId, err := r.database.CreateFoodItem(name, category, amount, householdId)
		if err != nil {
			return nil, err
		}
		return newFoodItemId, nil
	}
	err := fmt.Errorf("type-checking error: a combination of name, category, amount, and householdId was misformed")
	log.Print(err)
	return nil, err
}

func (r *Resolver) DeleteFoodItemMutationResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	var err error
	if ok {
		err = r.database.DeleteFoodItem(id)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	err = fmt.Errorf("type-checking error: id was not an int")
	log.Print(err)
	return nil, err
}

func (r *Resolver) UpdateFoodItemMutationResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	var err error
	if ok {
		err = r.database.UpdateFoodItem(id, p)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	err = fmt.Errorf("type-checking error: id was not an int")
	log.Print(err)
	return nil, err
}
