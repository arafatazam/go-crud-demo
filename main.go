package main

import (
  "log"
  "os"

  "github.com/joho/godotenv"
  _ "github.com/lib/pq"
  "github.com/golang-migrate/migrate"
  "github.com/golang-migrate/migrate/database/postgres"
  _ "github.com/golang-migrate/migrate/source/file"
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
  driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
  if err != nil {
	  log.Fatal("Migration db driver Failure")
  }
  m, err := migrate.NewWithDatabaseInstance(
      "file://migrations",
      "postgres", driver)
  if err != nil {
	  log.Fatal(err)
  }
  m.Up()

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