package main

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego"
	"github.com/forChin/test-crud/beego-crud/controllers"
	"github.com/forChin/test-crud/beego-crud/repository"

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

	beego.Router("/", &controllers.UserController{Repo: repo}, "post:SaveUser")
	beego.Run()
}
