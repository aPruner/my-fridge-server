package db

type User struct {
	ID       int
	Username string
}

type FoodItem struct {
	ID          int
	Name        string
	Category    string
	Amount      int
	HouseholdId int `sql:"household_id"`
	UserId      int `sql:"user_id"`
}

type Household struct {
	ID   int
	Name string
	City string
}
