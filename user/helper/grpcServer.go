package helper

import (
	"context"
	"github.com/lizhanfei/easygolib/server/grpc"
	"user/conf"
)

func InitGrpcServer(s *grpc.Server) {
	Logger.Infof(context.Background(), "init rpc")
	lis, err := grpc.NewListenerTcp(conf.GrpcConf.Address)
	if err != nil {
		panic("grpc listen fail")
	}
	s.SetListener(lis)
	s.Run()
}
