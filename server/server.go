package server

import (
	"github.com/aPruner/my-fridge-server/gql"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	*mux.Router
}

func Create() *Server {
	muxRouter := mux.NewRouter()
	muxRouter.Handle("/graphql", http.HandlerFunc(gql.GraphQLRequestHandler)).Methods("GET", "POST")
	return &Server{muxRouter}
}
