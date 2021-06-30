package logic

import (
	"context"
	"errors"
	"time"

	"userSystem/service/user/model"
	"userSystem/service/user/rpc/internal/svc"
	"userSystem/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateRequest) (*user.Response, error) {
	// todo: add your logic here and delete this line
	usersInfo, err := l.svcCtx.UserModel.FindManyByNameOrId(0, in.Name)
	if err != nil {
		return nil, errors.New("查询用户错误")
	}
	errorField := make(map[string]string)
	if len(usersInfo) != 0 {
		errorField["name"] = "用户名重复"
		return &user.Response{Id: 0, Name: "", ErrorField: errorField}, nil
	}
	created := time.Now()
	newUser, err := l.svcCtx.UserModel.Insert(model.Users{Id: 0, Name: in.Name, Password: in.Password, Created: created})
	if err != nil {
		return nil, errors.New("新建用户错误")
	}
	newUserId, _ := newUser.LastInsertId()

	return &user.Response{Id: newUserId, Name: in.Name, ErrorField: errorField}, nil
}
