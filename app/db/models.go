package db

type User struct {
	ID       int
	Username string
}

type FoodItem struct {
	ID             int
	Name           string
	Category       string
	Amount         int
	Unit           string
	HouseholdId    int `sql:"household_id"`
	ShoppingListId int `sql:"shopping_list_id"`
}

type Household struct {
	ID   int
	Name string
	City string
}

type ShoppingList struct {
	ID          int
	Name        string
	UserId      int    `sql:"user_id"`
	HouseholdId int    `sql:"household_id"`
	CreatedAt   string `sql:"created_at"`
}
