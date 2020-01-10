package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	pb "github.com/nonfu/laracom/demo-service/proto/demo"
	"log"
)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://laracom-consul:8500",
		}
	})
	service := micro.NewService(
		micro.Name("laracom.demo.cli"),
		micro.Registry(reg),
	)
	service.Init()

	client := pb.NewDemoServiceClient("laracom.service.demo", service.Client())
	rsp, err := client.SayHello(context.TODO(), &pb.DemoRequest{Name: "学院君"})
	if err != nil {
		log.Fatalf("服务调用失败：%v", err)
		return
	}
	log.Println(rsp.Text)
}