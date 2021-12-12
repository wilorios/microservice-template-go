//Package controller provides functions to translate incoming request
//into outcoming response
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wilorios/microservice-template-go/internal/entity"
	"github.com/wilorios/microservice-template-go/internal/service"
)

type PersonController interface {
	Save(ctx *gin.Context) entity.Person
	FindAll() []entity.Person
}

type personController struct {
	personService service.PersonService
}

//New function create a new controller
func New(ser service.PersonService) PersonController {
	return &personController{
		personService: ser,
	}
}

//FindAll function search all the entity person calling the service
func (c *personController) FindAll() []entity.Person {
	return c.personService.FindAll()
}

//Save function get the request and bind to entity and call the service
//to save the data
func (c *personController) Save(ctx *gin.Context) entity.Person {
	var e entity.Person
	ctx.BindJSON(&e)
	c.personService.Save(e)
	return e
}
