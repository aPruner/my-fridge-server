package gql

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"net/http"
)

type reqBody struct {
	GqlQuery string `json:"query"`
}

func GraphQLHandler(schema *graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Please provide a non-empty GQL query in the request body", 400)
			return
		}

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body, it was probably misformed", 400)
			return
		}

		gqlResult := ExecuteGraphQLQuery(rBody.GqlQuery, *schema)
		err = json.NewEncoder(w).Encode(gqlResult)
		if err != nil {
			http.Error(w, "Error encoding the response, our bad!", 500)
			return
		}
	}
}
