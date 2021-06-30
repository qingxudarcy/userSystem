package logic

import (
	"context"
	"errors"

	"userSystem/common/errorx"
	"userSystem/service/user/api/internal/svc"
	"userSystem/service/user/api/internal/types"
	"userSystem/service/user/rpc/userclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateLogic {
	return CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req types.CreateReq) (*types.ListReq, error) {
	// todo: add your logic here and delete this line
	rpcRes, err := l.svcCtx.UserRpc.CreateUser(l.ctx, &userclient.CreateRequest{Name: req.Name, Password: req.Password})
	if err != nil {
		return nil, errors.New("系统错误")
	}
	if len(rpcRes.ErrorField) != 0 {
		errorFields := make([]string, 0, len(rpcRes.ErrorField))
		for errorField := range rpcRes.ErrorField {
			errorFields = append(errorFields, errorField)
		}
		return nil, errorx.NewCheckError(errorFields[0], rpcRes.ErrorField[errorFields[0]])
	}

	return &types.ListReq{Id: rpcRes.Id, Name: req.Name}, nil
}
