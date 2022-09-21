package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lizhanfei/easygolib/server/grpc"
	"github.com/spf13/cobra"
	"sync"
	"user/app/router"
	"user/env"
	"user/helper"
)

func main() {
	engine := gin.New()
	ctxBase := context.Background()
	env.InitPath()
	helper.Init()

	var rootCmd = &cobra.Command{
		Use:   "user",
		Short: "user application ",
		Run: func(cmd *cobra.Command, args []string) {
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				httpServer(engine)
			}()
			go func() {
				defer wg.Done()
				grpcServer(ctxBase)
			}()
			wg.Wait()
		},
	}
	router.Commands(rootCmd, engine)

	if err := rootCmd.Execute(); err != nil {
		panic(err.Error())
	}
}

func httpServer(engine *gin.Engine) {
	router.RegisterRouter(engine)
	helper.InitHttpServer(engine)
}

func grpcServer(ctxBase context.Context) {
	grpcServer := grpc.NewServer(ctxBase)
	router.InitGrpc(grpcServer.Server)
	helper.InitGrpcServer(grpcServer)
}