package gql

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
)

func GraphQLRequestHandler(w http.ResponseWriter, r *http.Request)  {
	// TODO: Pipe status code through to logging middleware
	bytesWritten, err := w.Write([]byte(`hello worlds`))
	if err != nil {
		log.Fatalf("There was an error writing the response")
	}
	log.Printf("Request received, wrote %d bytes to response", bytesWritten)
}

// TODO: Get rid of this gql boilerplate
func Test() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			hello
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {"data":{"hello":"world"}}
}
