package main

import (
  "log"
  "os"

  "github.com/joho/godotenv"
  _ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"
  "github.com/gofiber/fiber"
  "github.com/gofiber/fiber/middleware"
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
	  log.Fatal("Database Connection Failure")
  }
  usrSvc := &UserService{
    DB: *db,
  }
  validate := validator.New()

  app := fiber.New()
  app.Use(middleware.Recover())
  app.Use(middleware.Logger())

  SetHandlers(app, usrSvc, validate)
  
  app.Listen(os.Getenv("PORT"))
}