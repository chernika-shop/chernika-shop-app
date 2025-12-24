package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRestRouter(r *gin.Engine ) {
	// REST API
	api := r.Group("/api")
	{
		api.GET("/health", healthCheck)
	}

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

