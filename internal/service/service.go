//Package service provides code modularity,
//it has the business logic and rules which in turn calls the entities
package service

import "github.com/wilorios/microservice-template-go/internal/entity"

type iPerson interface {
	Save(entity.Person) entity.Person
	FindAll() []entity.Person
}

type PersonServ struct {
	entities []entity.Person
}

func New() *PersonServ {
	return &PersonServ{}
}

func (ser *PersonServ) Save(ent entity.Person) entity.Person {
	ser.entities = append(ser.entities, ent)
	return ent
}

func (ser *PersonServ) FindAll() []entity.Person {
	if ser.entities == nil {
		ser.entities = make([]entity.Person, 0)
	}
	return ser.entities
}
