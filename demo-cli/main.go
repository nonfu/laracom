package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/nonfu/laracom/common/wrapper/breaker/hystrix"
	pb "github.com/nonfu/laracom/demo-service/proto/demo"
	"log"
	"os"
	"github.com/nonfu/laracom/demo-service/trace"
)

func main() {
	hystrix.Configure([]string{"laracom.service.demo.DemoService.SayHello"})

	// 初始化追踪器
	t, io, err := tracer.NewTracer("laracom.demo.cli", os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()

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