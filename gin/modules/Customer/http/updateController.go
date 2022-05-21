package customer

import "github.com/gin-gonic/gin"

func UpdateController(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"id":  id,
		"err": nil,
	})
}
