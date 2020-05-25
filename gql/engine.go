package gql

import (
	"github.com/graphql-go/graphql"
	"log"
)

func ExecuteGraphQLQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Printf("There were errors in the GQL query: %v", result.Errors)
	}
	return result
}
