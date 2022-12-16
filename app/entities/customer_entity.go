package entities

import "time"

type CustomerEntity struct {
	ID          string    `json:"id" gorm:"primarykey"`
	Name        string    `json:"name" `
	DateCreated time.Time `gorm:"date_created" json:"date_created"`
}
