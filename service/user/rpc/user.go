package main

import (
	"flag"
	"fmt"

	LogConf "userSystem/logs/conf"
	"userSystem/service/user/rpc/internal/config"
	"userSystem/service/user/rpc/internal/server"
	"userSystem/service/user/rpc/internal/svc"
	"userSystem/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	c.RpcServerConf.Log = LogConf.DefaultLog
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
