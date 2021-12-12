package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wilorios/microservice-template-go/internal/configurations"
	"github.com/wilorios/microservice-template-go/internal/controller"
	"github.com/wilorios/microservice-template-go/internal/service"
)

var (
	xService    service.XService       = service.New()
	xController controller.XController = controller.New(xService)
)

// startWebServer starts the web server.
func SetupRouter(conf configurations.Application) {
	log.Println("msg", "starting http server", "http:", conf.Port)

	server := gin.Default()

	server.GET("/microservice", func(ctx *gin.Context) {
		ctx.JSON(200, xController.FindAll())
	})

	server.POST("/microservice", func(ctx *gin.Context) {
		ctx.JSON(200, xController.Save(ctx))
	})

	server.Run(":" + conf.Port)
	//return server
}
