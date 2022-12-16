package data

import (
	"macus/entities"
)

type DbProvider interface {
	Create(*entities.CustomerEntity) error
	Get() ([]entities.CustomerEntity, error)
	GetById(id string) (entities.CustomerEntity, error)
	Update(*entities.CustomerEntity) error
}
