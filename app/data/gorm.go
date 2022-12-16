package data

import (
	"macus/entities"

	"gorm.io/gorm"
)

// type DbProvider interface {
// 	Create(*entities.CustomerEntity) error
// 	Get() ([]entities.CustomerEntity, error)
// 	GetById(id string) (entities.CustomerEntity, error)
// 	Update(*entities.CustomerEntity) error
// }

type GormProvider struct {
	db *gorm.DB
}

func NewGormProvider(db *gorm.DB) *GormProvider {
	return &GormProvider{db: db}
}

func (s *GormProvider) Create(customer *entities.CustomerEntity) error {
	return s.db.Create(customer).Error
}

func (s *GormProvider) Get() ([]entities.CustomerEntity, error) {
	var customers []entities.CustomerEntity
	result := s.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (s *GormProvider) GetById(id string) (entities.CustomerEntity, error) {
	var customer entities.CustomerEntity
	result := s.db.First(&customer, "id = ?", id)
	if result.Error != nil {
		return entities.CustomerEntity{}, result.Error
	}
	return customer, nil
}

func (s *GormProvider) Update(customer *entities.CustomerEntity) error {
	return s.db.Save(&customer).Error

}
