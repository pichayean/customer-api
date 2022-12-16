package models

type NewCustomer struct {
	Name string `json:"name" binding:"required" example:"choo" maxLength:"255"`
}

type UpdateCustomer struct {
	Name *string `json:"name" binding:"required"`
}
