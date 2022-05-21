package customer

import "github.com/gin-gonic/gin"

func AddController(c *gin.Context)  {
	c.JSON(200, gin.H{
		"err": nil,
	})
}