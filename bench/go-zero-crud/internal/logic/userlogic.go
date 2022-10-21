package logic

import (
	"context"
	"fmt"
	"log"

	"github.com/forChin/test-crud/go-zero-crud/internal/logic/models"
	"github.com/forChin/test-crud/go-zero-crud/internal/svc"
	"github.com/forChin/test-crud/go-zero-crud/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserReq) (resp *types.UserResp, err error) {
	u := models.User{Name: req.Name, Age: req.Age}
	err = l.svcCtx.Repo.SaveUser(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- saved", u.Name, u.Age)

	return nil, nil
}
