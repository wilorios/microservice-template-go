package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wilorios/microservice-template-go/internal/entity"
	"github.com/wilorios/microservice-template-go/internal/service"
)

type XController interface {
	FindAll() []entity.XEntity
	Save(ctx *gin.Context) entity.XEntity
}

type controller struct {
	service service.XService
}

func New(ser service.XService) XController {
	return &controller{
		service: ser,
	}
}

func (c *controller) FindAll() []entity.XEntity {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.XEntity {
	var e entity.XEntity
	ctx.BindJSON(&e)
	c.service.Save(e)
	return e
}
