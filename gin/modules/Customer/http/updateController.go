package customer

import (
	"net/http"

	"github.com/DanielSuhett/go/gin/modules/Customer/domain/entities"
	"github.com/DanielSuhett/go/gin/modules/Customer/services"
	"github.com/gin-gonic/gin"
)

func UpdateController(c *gin.Context) {
	id := c.Param("id")

	var customer entities.CreateCustomer
	service, err := services.Run()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := c.ShouldBind(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = service.UpdateCustomer(id, customer.Name, customer.Email, customer.Address)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(200, nil)
}
