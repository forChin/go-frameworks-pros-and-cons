package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/forChin/test-crud/gin-crud/model"
	"github.com/forChin/test-crud/gin-crud/repository"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	// Create a new endpoint
	r.POST("/", func(c *gin.Context) {
		var u model.User
		if err := c.ShouldBind(&u); err != nil {
			log.Fatal(err)
		}
		err = repo.SaveUser(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("-- saved", u.Name, u.Age)
	})

	log.Fatal(r.Run(":8080"))
}
