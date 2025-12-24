package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

)


func SetupGraphQLRouter(r *gin.Engine,graphQLHandler *handler.Server){
	// GraphQL endpoints с аннотациями
	r.POST("/graphql", graphQLPost(graphQLHandler))
	r.GET("/graphql", graphQLGet(graphQLHandler))
	r.GET("/playground", graphQLPlayground())

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