package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/nonfu/laracom/common/wrapper/breaker/hystrix"
	pb "github.com/nonfu/laracom/demo-service/proto/demo"
	userpb "github.com/nonfu/laracom/user-service/proto/user"
	"log"
)

type DemoServiceHandler struct {

}

func (s *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	rsp.Text = "你好, " + req.Name
	return nil
}

func (s *DemoServiceHandler) SayHelloByUserId(ctx context.Context, req *pb.HelloRequest, rsp *pb.DemoResponse) error {
	// 使用断路器
	hystrix.Configure([]string{"laracom.service.user.UserService.GetById"})
	service := micro.NewService(
		micro.WrapClient(hystrix.NewClientWrapper()),
	)
	client := userpb.NewUserServiceClient("laracom.service.user", service.Client())
	resp, err := client.GetById(context.TODO(), &userpb.User{Id: req.Id})
	if err != nil {
		return err
	}
	rsp.Text = "你好, " + resp.User.Name
	return nil
}

func main()  {
	// 注册服务名必须和 demo.proto 中的 package 声明一致
	service := micro.NewService(
		micro.Name("laracom.service.demo"),
	)
	service.Init()

	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}