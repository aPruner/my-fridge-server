package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
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

//func (d *Db) GetNextIdValueForTable(tableName string) (int, error) {
//	// TODO: Convert this query to ORM code if possible, this is just a shortcut
//	// TODO: But how do I inject table names (without using a model) with the ORM lib?
//	var maxId int
//	_, err := d.Model().QueryOne(pg.Scan(&maxId), fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1", tableName))
//	if err != nil {
//		log.Print(fmt.Errorf("there was an error in the GetNextIdValueForTable query: %s", err))
//		return -1, err
//	}
//
//	return maxId, nil
//}
