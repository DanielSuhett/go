package app

import (
	"log"

	"github.com/DanielSuhett/go/gin/infra/http/routes"
	"github.com/gin-gonic/gin"
)

func Gin() {
	app := gin.Default()

	routes.Router(app)

	err := app.Run(":8080")

	if(err != nil){
		log.Fatal(err)
	}
}