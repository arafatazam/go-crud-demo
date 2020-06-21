package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/go-playground/validator/v10"
)

func SetHandlers(app *fiber.App, usrSvc *UserService, validate *validator.Validate)  {
	users := app.Group("/users")
	
	users.Get("/", func (c *fiber.Ctx) {
		list, err := usrSvc.ListUsers()
		if err != nil {
			log.Panic(err)
		}
		if err:=c.JSON(&list); err!=nil {
			log.Panic(err)
		}
	})

	users.Post("/", func (c *fiber.Ctx) {
		c.Accepts("application/json")
		u := new(User)
		if err := c.BodyParser(u); err != nil {
		  log.Panic(err)
		}
		if err := validate.Struct(u); err!=nil {
		  log.Panic(err)
		}
		usrSvc.CreateUser(u)
		if err := c.JSON(&u); err != nil {
		  log.Panic(err)
		}
	  })
}