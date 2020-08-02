package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
	"os"
)

type Server struct {
	*mux.Router
}

func Create(schema *graphql.Schema) *Server {
	muxRouter := mux.NewRouter()

	// Build handler
	gqlHandler := handler.New(&handler.Config{
		Schema: schema,
		Pretty: true,
		GraphiQL: true,
	})

	// Apply middleware
	h := ApplyMiddlewarePipeline(gqlHandler)
	muxRouter.Handle("/graphql", h).Methods("GET", "POST")
	return &Server{muxRouter}
}

// Applies gorilla/handlers middleware
func ApplyMiddlewarePipeline(handler http.Handler) http.Handler {
	middlewareHandler := handlers.LoggingHandler(os.Stdout, handler)
	middlewareHandler = handlers.ContentTypeHandler(middlewareHandler, "application/json")
	return middlewareHandler
}
