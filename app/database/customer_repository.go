package database

import (
	"time"

	"github.com/google/uuid"
)

type ICustomerRepository interface {
	GetCustomerByID(id string) (CustomerEntity, error)
	CreateNewCustomer(newCustomer CustomerEntity, now time.Time) (CustomerEntity, error)
	ListCustomer() ([]CustomerEntity, error)
	Update(id string, update CustomerEntity, now time.Time) error
}

type CustomerRepository struct {
	provider DbProvider
}

func NewCustomerRepository(p DbProvider) *CustomerRepository {
	return &CustomerRepository{provider: p}
}
func (repository *CustomerRepository) CreateNewCustomer(newCustomer CustomerEntity, now time.Time) (CustomerEntity, error) {
	newID := uuid.New()
	customer := CustomerEntity{
		ID:          newID.String(),
		Name:        newCustomer.Name,
		DateCreated: now.UTC(),
	}
	var p = repository.provider
	if err := p.Create(&customer); err != nil {
		return CustomerEntity{}, err
	}
	return customer, nil
}

func (repository *CustomerRepository) ListCustomer() ([]CustomerEntity, error) {
	var p = repository.provider
	customers, err := p.Get()
	if err != nil {
		return []CustomerEntity{}, err
	}
	return customers, nil
}

func (repository *CustomerRepository) GetCustomerByID(id string) (CustomerEntity, error) {
	var p = repository.provider
	customer, err := p.GetById(id)
	if err != nil {
		return CustomerEntity{}, err
	}
	return customer, nil
}

func (repository *CustomerRepository) Update(id string, update CustomerEntity, now time.Time) error {
	var p = repository.provider
	customer, err := p.GetById(id)
	if err != nil {
		return err
	}
	customer.Name = update.Name

	if err := p.Update(&customer); err != nil {
		return err
	}
	return nil
}
