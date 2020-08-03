package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/graphql-go/graphql"
	"log"
)

type Db struct {
	*pg.DB
}

func BuildDbOptions(port string, user string, password string, dbName string) pg.Options {
	return pg.Options{
		Addr:     port,
		User:     user,
		Password: password,
		Database: dbName,
	}
}

func Create(options pg.Options) (*Db, error) {
	db := pg.Connect(&options)
	if err := db.Ping(context.TODO()); err != nil {
		return nil, err
	}
	return &Db{db}, nil
}

func (d *Db) GetUsersByUsername(username string) ([]User, error) {
	var users []User
	err := d.Model(&users).Where("username = ?", username).Select()
	if err != nil {
		log.Print(fmt.Errorf("there was an error in the GetUsersByUsername query: %s", err))
		return users, err
	}
	return users, nil
}

func (d *Db) GetFoodItemsByHouseholdId(householdId int) ([]FoodItem, error) {
	var foodItems []FoodItem
	err := d.Model(&foodItems).Where("household_id = ?", householdId).Select()
	if err != nil {
		log.Print(fmt.Errorf("there was an error in the GetFoodItemsByHouseholdId query: %s", err))
		return foodItems, err
	}
	return foodItems, nil
}

func (d *Db) GetHouseholdIdByUserId(userId int) (int, error) {
	var household Household
	err := d.Model(&household).Where("userId = ?", userId).Select()
	if err != nil {
		log.Print(fmt.Errorf("there was an error in the GetHouseholdIdByUserId query: %s", err))
		// Return -1 here as correct household IDs are >= 0
		return -1, err
	}
	return household.ID, nil
}

func (d *Db) CreateFoodItem(name string, category string, amount int, householdId int) (int, error) {
	foodItem := &FoodItem{
		Name:        name,
		Category:    category,
		Amount:      amount,
		HouseholdId: householdId,
	}

	err := d.Insert(foodItem)
	if err != nil {
		log.Print(fmt.Errorf("there was an error in the CreateFoodItem query: %s", err))
		return -1, err
	}

	// Assuming all went well, return the Id of the new FoodItem
	return foodItem.ID, nil
}

func (d *Db) DeleteFoodItem(id int) error {
	foodItem := &FoodItem{
		ID: id,
	}
	err := d.Delete(foodItem)
	if err != nil {
		log.Print(fmt.Errorf("there was an error in the DeleteFoodItem query: %s", err))
		return err
	}
	return nil
}

func (d *Db) UpdateFoodItem(id int, p graphql.ResolveParams) error {
	// TODO: Figure out how to do optional arguments for the GQL mutations
	foodItem := &FoodItem{
		ID: id,
		Name: p.Args["name"].(string),
		Category: p.Args["category"].(string),
		Amount: p.Args["amount"].(int),
		HouseholdId: p.Args["householdId"].(int),
	}

	err := d.Update(foodItem)
	if err != nil {
		log.Print(fmt.Errorf("there was an error in the UpdateFoodItem query: %s", err))
		return err
	}
	return nil
}
