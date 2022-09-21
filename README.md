## github.com/lizhanfei/easygolib demo项目
## user服务
提供http接口服务和grpc服务。

### http api
- /user/user/add
- /user/user/rm
- /user/user/login
- /user/user/auth

### grpc 
- UserServiceImplGrpc.Auth()

### 项目目录结构
```shell script
./
├── app  --业务目录
│   ├── controller  --控制器目录
│   │   ├── dto.go
│   │   └── userController.go
│   ├── entity      --实体定义目录
│   │   ├── user.go
│   │   ├── user.pb.go
│   │   └── user.proto
│   ├── model       --model层目录
│   │   ├── userModel.go
│   │   └── userToken.go
│   ├── router      --路由目录
│   │   ├── command.go
│   │   ├── grpc.go
│   │   └── http.go
│   └── service     --service 目录
│       ├── userService.go
│       ├── userServiceImplGrpc.go
│       └── userServiceImplV1.go
├── conf --配置管理
│   ├── app  --配置文件
│   │   ├── db.yaml
│   │   ├── grpc.yaml
│   │   ├── log.yaml
│   │   └── server.yaml
│   ├── conf.go  --配置解析引导
│   ├── grpc.go  --grpc配置
│   ├── log.go   --log配置
│   ├── mysql.go --gorm配置
│   ├── redis.go --redis配置
│   └── server.go --http配置
├── env
│   └── env.go   --环境变量
├── go.mod
├── go.sum
├── helper
│   ├── grpcServer.go  --启动grpc 
│   ├── httpServer.go  --启动http
│   ├── init.go    --资源初始化引导
│   ├── log.go     --初始化日志
│   ├── mysql.go   --初始化gortm
├── log  -- 日志目录
│   ├── order.server.log  --服务日志
│   └── order.wf.log  -- 异常日志
├── main.go
└── sql
    └── user.sql  -- 项目依赖sql
```
