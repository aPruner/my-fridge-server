package db

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log"
)

type Db struct {
	*pg.DB
}

func BuildDbOptions(port string, user string, password string, dbName string) pg.Options {
	return pg.Options{
		Addr: port,
		User: user,
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

func (d *Db) GetUsersByUsername(username string) []User {
	var users []User
	err := d.Model(&users).Where("username = ?", username).Select()
	if err != nil {
		log.Printf("There were errors in the GetUsersByUsername query: %v", err)
	}
	return users
}

func (d *Db) GetFoodItemsByHousehold(householdId int) []FoodItem {
	var foodItems []FoodItem
	err := d.Model(&foodItems).Where("householdId = ?", householdId).Select()
	if err != nil {
		log.Printf("There were errors in the GetFoodItemsByHousehold query: #{err}")
	}
	return foodItems
}
