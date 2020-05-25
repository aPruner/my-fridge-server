package gql

import (
	"github.com/aPruner/my-fridge-server/db"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	database *db.Db
}

// UserResolver initiates db query to get a user
func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	// Type-check the name
	username, ok := p.Args["username"].(string)
	if ok {
		_ = username
		// TODO: Write this function and its respective query in the db package
		// users := r.database.GetUsersByUsername(username)
		return nil, nil // will need to return users, nil
	}

	return nil, nil
}

