package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wilorios/microservice-template-go/internal/controller"
	"github.com/wilorios/microservice-template-go/internal/service"
)

var (
	xService    service.XService       = service.New()
	xController controller.XController = controller.New(xService)
)

func main() {
	server := gin.Default()

	server.GET("/microservice", func(ctx *gin.Context) {
		ctx.JSON(200, xController.FindAll())
	})

	server.POST("/microservice", func(ctx *gin.Context) {
		ctx.JSON(200, xController.Save(ctx))
	})

	server.Run(":8080")
}
