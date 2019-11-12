package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "laracom/demo-service/proto/demo"
)

type DemoServiceHandler struct {

}

func (s *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	rsp = &pb.DemoResponse{Text: "你好, " + req.Name}
	return nil
}

func main()  {
	// 注册服务名必须和 demo.proto 中的 package 声明一致
	service := micro.NewService(micro.Name("laracom.demo.service"))
	service.Init();
	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{})
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}