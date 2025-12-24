package main

import (
	"log"
	"os"
	
	"github.com/99designs/gqlgen/graphql/handler"
	"backend/graph"
	"backend/internal/router"
	"backend/pkg/database"
	_ "backend/api/docs"
)

// @title           Chernika Shop API
// @version         1.0
// @description     API для магазина женской одежды "Chernika"
// @host            localhost:8000
func main() {
	log.Println("The backend is launched")
	if err := database.RunMigrations(); err != nil{
		log.Fatalf("Failed to run migrations:%v",err)
	}

	// GraphQL handler
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{}},
		),
	)
	r := router.SetupRouter(srv)
	port := os.Getenv("SERVER_PORT")
	r.Run(":" + port)
}

