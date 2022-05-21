package routes

import "github.com/gin-gonic/gin"

func Router(g *gin.Engine) *gin.RouterGroup {
	router := g.Group("/v1")

	Customer(router)

	return router
}
