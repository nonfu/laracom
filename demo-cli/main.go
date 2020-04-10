package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/nonfu/laracom/common/wrapper/breaker/hystrix"
	pb "github.com/nonfu/laracom/demo-service/proto/demo"
	"log"
)

func main() {
	hystrix.Configure([]string{"laracom.service.demo.DemoService.SayHello"})
	service := micro.NewService(
		micro.Name("laracom.demo.cli"),
		micro.WrapClient(hystrix.NewClientWrapper()),
	)
	service.Init()

	client := pb.NewDemoServiceClient("laracom.service.demo", service.Client())

	rsp, err := client.SayHelloByUserId(context.TODO(), &pb.HelloRequest{Id: "1"})
	if err != nil {
		log.Fatalf("服务调用失败：%v", err)
		return
	}
	log.Println(rsp.Text)
}