package main

import (
  "log"
  "os"

  "github.com/joho/godotenv"
  _ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"
  "github.com/gofiber/fiber"
  "github.com/go-playground/validator/v10"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  dbUrl := os.Getenv("DB_URL")
  db := sqlx.MustConnect("postgres", dbUrl)

  err = db.Ping()
  if err != nil {
	  log.Panic("Database Connection Failure")
  }
  usrSvs := UserService{
    DB: *db,
  }
  validate := validator.New()

  app := fiber.New()

  app.Post("/users", func (c *fiber.Ctx) {
    c.Accepts("application/json")
    u := new(User)
    if err := c.BodyParser(u); err != nil {
      log.Fatal(err)
    }
    err := validate.Struct(u)
    if err!=nil {
      log.Panic(err)
    }
    usrSvs.CreateUser(u)
    if err := c.JSON(&u); err != nil {
      c.Next(err)
    }
  })
  
  app.Listen(os.Getenv("PORT"))
}