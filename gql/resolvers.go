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
	name, ok := p.Args["name"].(string)
	if ok {
		_ = name
		// users := r.database.GetUsersByName(name)
		return nil, nil
	}

	return nil, nil
}

