package database

// import (
// 	"fmt"
// 	"strconv"
// 	"time"

//     "github.com/google/uuid"
// )

// type CustomerRepository interface {
// 	GetCustomerByID(id string) (CustomerEntity, error)
// 	CreateNewCustomer(newCustomer CustomerEntity, now time.Time) (CustomerEntity, error)
// 	ListCustomer() ([]CustomerEntity, error)
// 	Update(id string, update CustomerEntity, now time.Time) error
// }

// type DbProvider struct {
// 	DB *GormStore
// }

// func (provider DbProvider) CreateNewCustomer(newCustomer CustomerEntity, now time.Time) (CustomerEntity, error) {
//     newID := uuid.New()
// 	customer := CustomerEntity{
// 		ID:          newID.String(),
// 		Name:        newCustomer.Name,
// 		DateCreated: now.UTC(),
// 	}

// 	const query = `INSERT INTO Customers (customer_id, name, date_created)VALUES ($1, $2, $3)`
// 	tx := postgres.DB.MustBegin()
// 	tx.MustExec(query, customer.ID, customer.Name, customer.DateCreated)
// 	if err := tx.Commit(); err != nil {
// 		return CustomerEntity{}, err
// 	}
// 	return customer, nil
// }

// func (provider DbProvider) ListCustomer() ([]CustomerEntity, error) {
// 	var customer []CustomerEntity
// 	const query = `SELECT customer_id, name, price, amount, date_created, date_updated FROM Customers`
// 	err := postgres.DB.Select(&customer, query)
// 	if err != nil {
// 		return []CustomerEntity{}, err
// 	}
// 	for index, prod := range customer {
// 		customer[index].DateCreated = prod.DateCreated.UTC()
// 	}
// 	return customer, nil
// }

// func (provider DbProvider) GetCustomerByID(id string) (CustomerEntity, error) {
// 	var customer CustomerEntity
// 	const query = `SELECT customer_id, name, date_created FROM Customers WHERE customer_id=$1`
// 	err := postgres.DB.Get(&customer, query, id)
// 	if err != nil {
// 		return CustomerEntity{}, err
// 	}
// 	customer.DateCreated = customer.DateCreated.UTC()
// 	return customer, nil
// }

// func (provider DbProvider) Update(id string, update CustomerEntity, now time.Time) error {
// 	const query = `UPDATE Customers SET "name" = $2 WHERE customer_id=$1`
// 	tx := postgres.DB.MustBegin()
// 	tx.MustExec(query, update.ID, update.Name)
// 	if err := tx.Commit(); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (provider DbProvider) getNewID() (ID string) {
// 	var customers []CustomerEntity
// 	const query1 = `SELECT customer_id,name, date_created FROM Customers ORDER BY customer_id DESC`
// 	err := postgres.DB.Select(&customers, query1)
// 	var newID = "0"
// 	if err != nil {
// 		newID = "0"
// 	}
// 	newID = customers[0].ID
// 	intVar, errint := strconv.Atoi(newID)
// 	if errint != nil {
// 		fmt.Println(errint)
// 	}
// 	return strconv.Itoa(intVar + 1)
// }
