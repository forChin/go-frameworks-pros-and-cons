package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/forChin/test-crud/httprouter-crud/model"
	"github.com/forChin/test-crud/httprouter-crud/repository"
	"github.com/julienschmidt/httprouter"
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

	r := httprouter.New()
	r.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var u model.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			log.Fatal(err)
		}
		err = repo.SaveUser(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("-- saved", u.Name, u.Age)
	})

	log.Fatal(http.ListenAndServe(":8080", r))

}
