package main

import (
	"os"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"backend/graph"
	_ "backend/api/docs"
)	

// @title           Chernika Shop API
// @version         1.0
// @description     API для магазина женской одежды "Chernika"
// @host            localhost:8000
// @BasePath        /api
func main() {
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// REST API
	api := r.Group("/api")
	{
		api.GET("/health", healthCheck)
	}

	// GraphQL handler
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{}},
		),
	)

	// GraphQL endpoints с аннотациями
	r.POST("/graphql", graphQLPost(srv))
	r.GET("/graphql", graphQLGet(srv))
	
	// GraphQL Playground
	r.GET("/playground", graphQLPlayground())

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}
	r.Run(":" + port)
}

// @Summary     Health check
// @Description Проверка работоспособности API
// @Tags        system
// @Accept      json
// @Produce     json
// @Success     200  {object}  map[string]string
// @Router      /api/health [get]
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

// @Summary     GraphQL Query/Mutation
// @Description Выполнить GraphQL запрос или мутацию
// @Tags        graphql
// @Accept      json
// @Produce     json
// @Param       query body object true "GraphQL запрос" example({"query":"query { user(id: 1) { name email } }"})
// @Success     200  {object}  object
// @Failure     400  {object}  object
// @Router      /graphql [post]
func graphQLPost(srv *handler.Server) gin.HandlerFunc {
	return gin.WrapH(srv)
}

// @Summary     GraphQL Query (GET)
// @Description Выполнить GraphQL запрос через GET
// @Tags        graphql
// @Accept      json
// @Produce     json
// @Param       query query string true "GraphQL запрос" example("query { user(id: 1) { name } }")
// @Success     200  {object}  object
// @Router      /graphql [get]
func graphQLGet(srv *handler.Server) gin.HandlerFunc {
	return gin.WrapH(srv)
}

// @Summary     GraphQL Playground
// @Description Интерактивная среда для тестирования GraphQL запросов
// @Tags        graphql
// @Produce     text/html
// @Success     200
// @Router      /playground [get]
func graphQLPlayground() gin.HandlerFunc {
	return gin.WrapH(playground.Handler("GraphQL Playground", "/graphql"))
}