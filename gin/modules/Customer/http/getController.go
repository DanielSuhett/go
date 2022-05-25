package customer

import (
	"net/http"

	"github.com/DanielSuhett/go/gin/modules/Customer/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetController(c *gin.Context) {
	id := c.Param("id")

	service, err := services.Run()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	customer, err := service.GetCustomer(id)

	if err == mongo.ErrNoDocuments {
		c.Status(404)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{
		"customer": customer,
	})
}
