package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
)

type UserService struct{
	DB sqlx.DB
}

func (us UserService) CreateUser(user *User) error{
	user.Id = uuid.New().String()
	return nil
}