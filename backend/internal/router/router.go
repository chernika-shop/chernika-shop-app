package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(graphQLHandler	*handler.Server) *gin.Engine{
	r := gin.Default()

	setupSwagger(r)

	SetupRestRouter(r)

	SetupGraphQLRouter(r,graphQLHandler)

	return r
}

func setupSwagger(r *gin.Engine) {
	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})
	swaggerHandler := ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.DocExpansion("list"),
		ginSwagger.DeepLinking(true),
	)
	r.GET("/swagger/*any", swaggerHandler)
}