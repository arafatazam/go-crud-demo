package main

import (
	"log"

	"github.com/google/uuid"
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
		u := User{}
		if err := c.BodyParser(&u); err != nil {
		  log.Panic(err)
		}
		if err := validate.Struct(u); err!=nil {
		  log.Panic(err)
		}
		usrSvc.CreateUser(&u)
		if err := c.JSON(&u); err != nil {
		  log.Panic(err)
		}
	  })

	users.Get("/:id", func(c *fiber.Ctx){
		id := c.Params("id")
		if _, err:= uuid.Parse(id); err!=nil{
			log.Panic("Invalid UUID")
		}
		u, err := usrSvc.GetUser(id)
		if err!=nil {
			log.Panic(err)
		}
		c.JSON(&u)
	})

	users.Put("/:id", func(c *fiber.Ctx){
		c.Accepts("application/json")
		id := c.Params("id")
		if _, err:= uuid.Parse(id); err!=nil{
			c.Next(fiber.NewError(400, "The given uuid cannot be parsed."))
			return
		}
		u := User{}
		if err := c.BodyParser(&u); err != nil {
		  log.Panic(err)
		}
		if err := validate.Struct(u); err!=nil {
		  log.Panic(err)
		}
		usrSvc.UpdateUser(id, &u)
		c.JSON(&u)
	})

	users.Delete("/:id", func(c *fiber.Ctx){
		id := c.Params("id")
		if _, err:= uuid.Parse(id); err!=nil{
			c.Next(fiber.NewError(400, "The given uuid cannot be parsed."))
			return
		}
		usrSvc.DeleteUser(id)
		c.Send("Deleted")
	})
}