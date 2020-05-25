package server

import (
	"github.com/aPruner/my-fridge-server/gql"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
)

type Server struct {
	*mux.Router
}

func Create(schema *graphql.Schema) *Server {
	muxRouter := mux.NewRouter()
	muxRouter.Handle("/graphql", gql.GraphQLHandler(schema)).Methods("GET", "POST")
	return &Server{muxRouter}
}
