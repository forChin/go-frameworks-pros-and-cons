package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/beego/beego"

	"github.com/forChin/test-crud/beego-crud/models"
	"github.com/forChin/test-crud/beego-crud/repository"
)

type UserController struct {
	beego.Controller
	Repo *repository.Repository
}

// Example:
//
//   req: GET /task/
//   res: 200 {"Tasks": [
//          {"ID": 1, "Title": "Learn Go", "Done": false},
//          {"ID": 2, "Title": "Buy bread", "Done": true}
//        ]}
func (this *UserController) SaveUser() {
	var u models.User
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &u); err != nil {
		log.Fatal(err)
	}
	err := this.Repo.SaveUser(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- saved", u.Name, u.Age)
}
