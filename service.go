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

func (svc UserService) GetUser(id string) (User, error){
	u := User{}
	if err:= svc.DB.Get(&u, svc.DB.Rebind(`SELECT * FROM "users" WHERE "id" = ? LIMIT 1`), id); err!=nil {
		return u, err
	}
	return u, nil
}

func (svc UserService) UpdateUser(id string, u *User) {
	update := svc.DB.Rebind(`UPDATE "users" SET "email" = ?, "first_name" = ?, "last_name" = ?,
				"address" = ?, "phone" = ? WHERE "id" = ?`)
	svc.DB.MustExec(update, u.Email, u.FirstName, u.LastName, u.Address, u.Phone, id)
	u.Id = id
}

func (svc UserService) DeleteUser(id string) {
	update := svc.DB.Rebind(`DELETE FROM "users" WHERE (("id" = ?))`)
	svc.DB.MustExec(update, id)
}