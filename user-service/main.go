package main

import (
	"fmt"
	"github.com/micro/go-micro"
	database "github.com/nonfu/laracom/user-service/db"
	"github.com/nonfu/laracom/user-service/handler"
	pb "github.com/nonfu/laracom/user-service/proto/user"
	repository "github.com/nonfu/laracom/user-service/repo"
	"github.com/nonfu/laracom/user-service/service"
	"log"
)

func main() {

	// 创建数据库连接，程序退出时断开连接
	db, err := database.CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// 和 Laravel 数据库迁移类似
	// 每次启动服务时都会检查，如果数据表不存在则创建，已存在检查是否有修改
	db.AutoMigrate(&pb.User{})

	// 初始化 Repo 实例用于后续数据库操作
	repo := &repository.UserRepository{db}
	// 初始化 token service
	token := &service.TokenService{repo}

	// 以下是 Micro 创建微服务流程
	srv := micro.NewService(
		micro.Name("laracom.service.user"),
		micro.Version("latest"),  // 新增接口版本参数
	)
	srv.Init()


	// 注册处理器
	pb.RegisterUserServiceHandler(srv.Server(), &handler.UserService{repo, token})

	// 启动用户服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
