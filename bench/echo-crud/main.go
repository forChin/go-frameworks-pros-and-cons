package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/forChin/test-crud/echo-crud/model"
	"github.com/forChin/test-crud/echo-crud/repository"
	"github.com/labstack/echo/v4"
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

	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		var u model.User
		if err := c.Bind(&u); err != nil {
			log.Fatal(err)
		}
		err = repo.SaveUser(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("-- saved", u.Name, u.Age)

		return nil
	})
	e.Logger.Fatal(e.Start(":8080"))
}
