package data

import (
	"macus/entities"

	"gorm.io/gorm"
)

type GormProvider[T any] struct {
	db *gorm.DB
}

func NewGormProvider[T any](db *gorm.DB) *GormProvider[T] {
	return &GormProvider[T]{db: db}
}

func (s *GormProvider[T]) Create(customer *T) error {
	return s.db.Create(customer).Error
}

func (s *GormProvider[T]) Get() ([]T, error) {
	var entity []T
	result := s.db.Find(&entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

func (s *GormProvider[T]) GetById(id string) (T, error) {
	var entity T
	result := s.db.First(&entity, "id = ?", id)
	if result.Error != nil {
		return entity, result.Error
	}
	return entity, nil
}

func (s *GormProvider[T]) Update(customer *entities.CustomerEntity) error {
	return s.db.Save(&customer).Error

}
