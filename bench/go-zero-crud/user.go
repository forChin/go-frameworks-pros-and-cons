package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/forChin/test-crud/go-zero-crud/internal/config"
	"github.com/forChin/test-crud/go-zero-crud/internal/handler"
	"github.com/forChin/test-crud/go-zero-crud/internal/logic/repository"
	"github.com/forChin/test-crud/go-zero-crud/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	_ "github.com/lib/pq"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

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

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c, repo)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
