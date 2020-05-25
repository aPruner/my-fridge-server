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
