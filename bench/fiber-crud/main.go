package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/forChin/test-crud/fiber-crud/model"
	"github.com/forChin/test-crud/fiber-crud/repository"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "2000postgres"
	dbname := "sammy"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	repo := repository.New(db)

	app := fiber.New() // create a new Fiber instance
	// Create a new endpoint
	app.Post("/", func(c *fiber.Ctx) error {
		var u model.User
		if err := c.BodyParser(&u); err != nil {
			log.Fatal(err)
		}
		err = repo.SaveUser(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("-- saved", u.Name, u.Age)

		return nil
	})

	log.Fatal(app.Listen(":8080"))
}
