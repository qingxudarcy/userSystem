package logic

import (
	"context"
	"errors"
	"strings"
	"time"

	"userSystem/common/errorx"
	"userSystem/service/user/api/internal/svc"
	"userSystem/service/user/api/internal/types"
	"userSystem/service/user/model"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginResp, error) {
	// todo: add your logic here and delete this line
	userInfo, errFieldName, err := loginCheck(l, req.Name, req.Password)
	if err != "" {
		if errFieldName != "" {
			return nil, errorx.NewCheckError(errFieldName, err)
		} else {
			return nil, errors.New("系统错误,请联系管理员")
		}
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, tokenErr := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &types.LoginResp{
		Id:           userInfo.Id,
		Name:         userInfo.Name,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))

}

func loginCheck(l *LoginLogic, name, password string) (*model.Users, string, string) {
	if len(strings.TrimSpace(name)) == 0 {
		return nil, "name", "长度有误"
	}
	if len(strings.TrimSpace(password)) == 0 {
		return nil, "password", "长度有误"
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByName(name)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, "name", "用户名不存在"
	default:
		return nil, "", "错误"
	}
	if userInfo.Password != password {
		return nil, "password", "密码错误"
	}

	return userInfo, "", ""

}
