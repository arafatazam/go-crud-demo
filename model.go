package main

type User struct{
	Id string					`json:"id" validate:"omitempty,uuid4"`
	Email string				`json:"email" validate:"required,email,max=100"`
	FirstName string			`json:"first_name" validate:"required,alpha,max=50"`
	LastName string				`json:"second_name,omitempty" validate:"omitempty,alpha,max=50"`
	Address string				`json:"address" validate:"required,max=50"`
	Phone string				`json:"phone" validate:"required,numeric,max=20"`
}