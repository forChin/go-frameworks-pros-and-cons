package svc

import (
	"github.com/forChin/test-crud/go-zero-crud/internal/config"
	"github.com/forChin/test-crud/go-zero-crud/internal/logic/repository"
)

type ServiceContext struct {
	Config config.Config
	Repo   *repository.Repository
}

func NewServiceContext(c config.Config, repo *repository.Repository) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   repo,
	}
}
