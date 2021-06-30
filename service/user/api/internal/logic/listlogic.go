package logic

import (
	"context"
	"errors"
	"userSystem/service/user/api/internal/svc"
	"userSystem/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type ListResps []*types.ListResp

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListLogic {
	return ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req types.ListReq) (ListResps, error) {
	// todo: add your logic here and delete this line
	usersInfo, err := l.svcCtx.UserModel.FindManyByNameOrId(req.Id, req.Name)
	switch err {
	case nil:
	default:
		logx.Errorf("数据库操作错误: %s", err)
		return nil, errors.New("查询错误")
	}
	var items ListResps
	for _, user := range usersInfo {
		items = append(items, &types.ListResp{Id: user.Id, Name: user.Name})
	}

	return items, nil
}
