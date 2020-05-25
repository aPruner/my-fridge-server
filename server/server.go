package server

import (
	"github.com/aPruner/my-fridge-server/gql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"net/http"
	"os"
)

type Server struct {
	*mux.Router
}

func Create(schema *graphql.Schema) *Server {
	muxRouter := mux.NewRouter()

	// Apply middleware
	gqlHandler := ApplyMiddlewarePipeline(gql.GraphQLHandler(schema))
	muxRouter.Handle("/graphql", gqlHandler).Methods("GET", "POST")
	return &Server{muxRouter}
}

// Applies gorilla/handlers middleware
func ApplyMiddlewarePipeline(handler http.Handler) http.Handler {
	middlewareHandler := handlers.LoggingHandler(os.Stdout, handler)
	middlewareHandler = handlers.ContentTypeHandler(middlewareHandler, "application/json")
	return middlewareHandler
}
