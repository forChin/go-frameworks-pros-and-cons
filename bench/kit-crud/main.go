package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/forChin/test-crud/kit-crud/model"
	"github.com/forChin/test-crud/kit-crud/repository"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	_ "github.com/lib/pq"
)

// func main() {
// 	host := "localhost"
// 	port := 5432
// 	user := "postgres"
// 	password := "2000postgres"
// 	dbname := "sammy"

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}

// 	repo := repository.New(db)

// 	r := httprouter.New()
// 	r.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		var u model.User
// 		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
// 			log.Fatal(err)
// 		}
// 		err = repo.SaveUser(u)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("-- saved", u.Name, u.Age)
// 	})

// 	log.Fatal(http.ListenAndServe(":8080", r))

// }

type UserService interface {
	SaveUser(name string, age int)
}

type userService struct {
	repo *repository.Repository
}

func (us userService) SaveUser(name string, age int) {
	u := model.User{Name: name, Age: age}
	err := us.repo.SaveUser(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-- saved", u.Name, u.Age)
}

type saveUserRequest struct {
	Name string
	Age  int
}

type saveUserResponse struct {
}

func makeSaveUserEndpoint(uvc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(saveUserRequest)
		uvc.SaveUser(req.Name, req.Age)
		return saveUserResponse{}, nil
	}
}

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
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

	uvc := userService{repo}

	saveUserHandler := httptransport.NewServer(
		makeSaveUserEndpoint(uvc),
		decodeSaveUserRequest,
		encodeResponse,
	)

	http.Handle("/", saveUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeSaveUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request saveUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
