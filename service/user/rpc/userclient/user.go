// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

//go:generate mockgen -destination ./user_mock.go -package userclient -source $GOFILE

package userclient

import (
	"context"

	"userSystem/service/user/rpc/user"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Response      = user.Response
	UpdateRequest = user.UpdateRequest
	CreateRequest = user.CreateRequest

	User interface {
		CreateUser(ctx context.Context, in *CreateRequest) (*Response, error)
		UpdateUser(ctx context.Context, in *UpdateRequest) (*Response, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) CreateUser(ctx context.Context, in *CreateRequest) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in)
}

func (m *defaultUser) UpdateUser(ctx context.Context, in *UpdateRequest) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in)
}
