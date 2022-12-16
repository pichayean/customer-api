package services

import (
	"macus/data"
	"macus/entities"
	"time"

	"github.com/google/uuid"
)

type CustomerService interface {
	GetCustomerByID(id string) (entities.CustomerEntity, error)
	CreateNewCustomer(newCustomer entities.CustomerEntity, now time.Time) (entities.CustomerEntity, error)
	ListCustomer() ([]entities.CustomerEntity, error)
	Update(id string, update entities.CustomerEntity, now time.Time) error
}

type customerRepository struct {
	provider data.DbProvider
}

func NewCustomerRepository(p data.DbProvider) *customerRepository {
	return &customerRepository{provider: p}
}

func (repository *customerRepository) CreateNewCustomer(newCustomer entities.CustomerEntity, now time.Time) (entities.CustomerEntity, error) {
	newID := uuid.New()
	customer := entities.CustomerEntity{
		ID:          newID.String(),
		Name:        newCustomer.Name,
		DateCreated: now.UTC(),
	}
	var p = repository.provider
	if err := p.Create(&customer); err != nil {
		return entities.CustomerEntity{}, err
	}
	return customer, nil
}

func (repository *customerRepository) ListCustomer() ([]entities.CustomerEntity, error) {
	customers, err := repository.provider.Get()
	if err != nil {
		return []entities.CustomerEntity{}, err
	}
	return customers, nil
}

func (repository *customerRepository) GetCustomerByID(id string) (entities.CustomerEntity, error) {
	customer, err := repository.provider.GetById(id)
	if err != nil {
		return entities.CustomerEntity{}, err
	}
	return customer, nil
}

func (repository *customerRepository) Update(id string, update entities.CustomerEntity, now time.Time) error {
	customer, err := repository.provider.GetById(id)
	if err != nil {
		return err
	}
	customer.Name = update.Name

	if err := repository.provider.Update(&customer); err != nil {
		return err
	}
	return nil
}
