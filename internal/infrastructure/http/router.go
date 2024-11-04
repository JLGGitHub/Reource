package http

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Grupo de rutas para usuarios
	chairGroup := router.Group("/resource/chair")
	{
		chairGroup.GET("/", func(c *gin.Context) { c.String(200, "Get all chairs") })
	}

	return router

}
