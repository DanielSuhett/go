package customer

import "github.com/gin-gonic/gin"

func GetController(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"id":  id,
		"err": nil,
	})
}
