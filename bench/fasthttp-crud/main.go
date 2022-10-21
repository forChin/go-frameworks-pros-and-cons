package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/forChin/test-crud/fasthttp-crud/model"
	"github.com/forChin/test-crud/fasthttp-crud/repository"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
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

	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			var u model.User
			if err := json.Unmarshal(ctx.PostBody(), &u); err != nil {
				log.Fatal(err)
			}
			err = repo.SaveUser(u)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("-- saved", u.Name, u.Age)

		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	log.Fatal(fasthttp.ListenAndServe(":8080", m))
}
