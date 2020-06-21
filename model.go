package main

type User struct{
	Id string					`json:"id" validate:"omitempty,uuid4"`
	Email string				`json:"email" validate:"required,email,max=100"`
	FirstName string			`json:"first_name" validate:"required,alpha,max=50" db:"first_name"`
	LastName string				`json:"last_name,omitempty" validate:"omitempty,alpha,max=50" db:"last_name"`
	Address string				`json:"address" validate:"required,max=50"`
	Phone string				`json:"phone" validate:"required,numeric,max=20"`
}