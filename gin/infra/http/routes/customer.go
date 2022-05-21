package routes

import (
	c "github.com/DanielSuhett/go/gin/modules/Customer/http"
	"github.com/gin-gonic/gin"
)

func Customer(r *gin.RouterGroup) {
	customer := r.Group("/customer")
	{
		customer.POST("/", c.AddController)
		customer.PUT("/:id", c.UpdateController)
		customer.GET("/:id", c.GetController)
	}
}
