package router

import (
	"google.golang.org/grpc"
	"user/app/entity"
	"user/app/service"
)

func InitGrpc(s *grpc.Server) {
	entity.RegisterUserServiceServer(s, &service.UserServiceImplGrpc{})
}
