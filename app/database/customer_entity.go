package database

import "time"

// Product is a item is stroe
//
//	type CustomerEntity struct {
//		ID          string    `db:"customer_id" json:"id"`
//		Name        string    `db:"name" json:"name"`
//		DateCreated time.Time `db:"date_created" json:"date_created"`
//	}
type CustomerEntity struct {
	ID          string    `json:"id" gorm:"primarykey"`
	Name        string    `json:"name" `
	DateCreated time.Time `gorm:"date_created" json:"date_created"`
}
