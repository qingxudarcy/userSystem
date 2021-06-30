package logic

import (
	"context"
	"errors"

	"userSystem/service/user/model"
	"userSystem/service/user/rpc/internal/svc"
	"userSystem/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateRequest) (*user.Response, error) {
	// todo: add your logic here and delete this line
	errorField := make(map[string]string)
	oldUserInfo, err := l.svcCtx.UserModel.FindOne(in.Id)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.Errorf("数据库操作错误: %s", err)
		errorField["id"] = "用户ID不存在"
		return &user.Response{Id: 0, Name: "", ErrorField: errorField}, nil
	default:
		logx.Errorf("数据库操作错误: %s", err)
		return nil, errors.New("系统错误")
	}

	existsUserInfo, err := l.svcCtx.UserModel.FindOneByName(in.Name)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
	default:
		logx.Errorf("数据库操作错误: %s", err)
		return nil, errors.New("系统错误")
	}
	if existsUserInfo != nil && existsUserInfo.Id != in.Id {
		errorField["name"] = "用户名重复"
		return &user.Response{Id: 0, Name: "", ErrorField: errorField}, nil
	}

	updateErr := l.svcCtx.UserModel.Update(model.Users{Name: in.Name, Password: oldUserInfo.Password, Id: oldUserInfo.Id, Created: oldUserInfo.Created})
	if updateErr != nil {
		logx.Errorf("数据库操作错误: %s", err)
		return nil, errors.New("系统错误")
	}

	return &user.Response{Name: in.Name, Id: in.Id, ErrorField: errorField}, nil
}
