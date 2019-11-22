package main

import (
	"context"
	"github.com/micro/go-micro"
	pb "github.com/nonfu/laracom/demo-service/proto/demo"
	"log"
)

type DemoServiceHandler struct {

}

func (s *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	rsp.Text = "你好, " + req.Name
	return nil
}

func main()  {
	// 注册服务名必须和 demo.proto 中的 package 声明一致
	service := micro.NewService(micro.Name("laracom.service.demo"))
	service.Init()

	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}