//Package service provides code modularity, 
//it has the business logic and rules which in turn calls the entities
package service

import "github.com/wilorios/microservice-template-go/internal/entity"

type PersonService interface {
	Save(entity.Person) entity.Person
	FindAll() []entity.Person
}

type personService struct {
	entities []entity.Person
}

func New() PersonService {
	return &personService{
		entities: []entity.Person{},
	}
}

func (ser *personService) Save(ent entity.Person) entity.Person {
	ser.entities = append(ser.entities, ent)
	return ent
}

func (ser *personService) FindAll() []entity.Person {
	return ser.entities
}
