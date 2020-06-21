package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
)

type UserService struct{
	DB sqlx.DB
}

func (svc UserService) ListUsers() ([]User, error) {
	uu := []User{}
	if err:= svc.DB.Select(&uu, `SELECT * FROM "users"`); err!=nil {
		return nil, err
	}
	return uu, nil
}

func (svc UserService) CreateUser(user *User) {
	user.Id = uuid.New().String()
	createUser := svc.DB.Rebind(`INSERT INTO "users" ("id", "email", "first_name", "last_name",
				 "address", "phone") VALUES (?, ?, ?, ?, ?, ?)`)
	svc.DB.MustExec(createUser, user.Id, user.Email, user.FirstName, 
		user.LastName, user.Address, user.Phone)
}