package main

import (
    "os"
    "github.com/99designs/gqlgen/graphql/handler"
    "backend/graph"
    "backend/internal/router"
    _ "backend/api/docs"
)

// @title           Chernika Shop API
// @version         1.0
// @description     API для магазина женской одежды "Chernika"
// @host            localhost:8000
func main() {
	// GraphQL handler
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{}},
		),
	)
	r := router.SetupRouter(srv)
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}

	r.Run(":" + port)
}

