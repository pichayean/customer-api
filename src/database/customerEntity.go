package database

import "time"

// Product is a item is stroe
type CustomerEntity struct {
	ID          string    `db:"customer_id" json:"id"`
	Name        string    `db:"name" json:"name"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
}
