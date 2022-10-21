package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/forChin/test-crud/kratos-crud/model"
	"github.com/forChin/test-crud/kratos-crud/repository"
	_ "github.com/lib/pq"

	"github.com/go-kratos/kratos/v2"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
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

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var u model.User
		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			log.Fatal(err)
		}

		err = repo.SaveUser(u)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("-- saved", u.Name, u.Age)

	}).Methods("POST")

	httpSrv := transhttp.NewServer(transhttp.Address(":8080"))
	httpSrv.HandlePrefix("/", router)

	app := kratos.New(
		kratos.Name("mux"),
		kratos.Server(
			httpSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
