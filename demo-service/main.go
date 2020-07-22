package main

import (
	"context"
	"github.com/micro/go-micro"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	pb "github.com/nonfu/laracom/demo-service/proto/demo"
	userpb "github.com/nonfu/laracom/user-service/proto/user"
	"log"
)

type DemoServiceHandler struct {

}

func (s *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	// 如果参数为空设置默认值
	if req.Name == "" {
		req.Name = "学院君"
	}
	rsp.Text = "你好, " + req.Name
	return nil
}

func (s *DemoServiceHandler) SayHelloByUserId(ctx context.Context, req *pb.HelloRequest, rsp *pb.DemoResponse) error {

	service := k8s.NewService(
		micro.Name("laracom.service.demo"),
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

	service := micro.NewService(
		micro.Name("laracom.service.demo"),
	)
	service.Init()

	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}