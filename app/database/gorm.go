package database

import (
	"gorm.io/gorm"
)

type DbProvider interface {
	Create(*CustomerEntity) error
	Get() ([]CustomerEntity, error)
	GetById(id string) (CustomerEntity, error)
	Update(*CustomerEntity) error
}

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{db: db}
}

func (s *GormStore) Create(customer *CustomerEntity) error {
	return s.db.Create(customer).Error
}

func (s *GormStore) Get() ([]CustomerEntity, error) {
	var customers []CustomerEntity
	result := s.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (s *GormStore) GetById(id string) (CustomerEntity, error) {
	var customer CustomerEntity
	result := s.db.First(&customer, "id = ?", id)
	if result.Error != nil {
		return CustomerEntity{}, result.Error
	}
	return customer, nil
}

func (s *GormStore) Update(customer *CustomerEntity) error {
	return s.db.Save(&customer).Error
	// return s.db.Model(&customer).Select("*").Update(CustomerEntity{Name: "jinzhu", Role: "admin", Age: 0}).Error

}
