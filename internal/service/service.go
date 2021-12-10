package service

import "github.com/wilorios/microservice-template-go/internal/entity"

type XService interface {
	Save(entity.XEntity) entity.XEntity
	FindAll() []entity.XEntity
}

type xService struct {
	entities []entity.XEntity
}

func New() XService {
	return &xService{
		entities: []entity.XEntity{},
	}
}

func (ser *xService) Save(ent entity.XEntity) entity.XEntity {
	ser.entities = append(ser.entities, ent)
	return ent
}

func (ser *xService) FindAll() []entity.XEntity {
	return ser.entities
}
