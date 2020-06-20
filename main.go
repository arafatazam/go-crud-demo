package main

import (
  "log"
  "os"

  "github.com/joho/godotenv"
  _ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"
  "github.com/gofiber/fiber"
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

  app := fiber.New()

  app.Post("/users", func (c *fiber.Ctx) {
    c.Send("Create User")
  })
  
  app.Listen(os.Getenv("PORT"))
}