package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/nonfu/laracom/common/wrapper/breaker/hystrix"
	pb "github.com/nonfu/laracom/demo-service/proto/demo"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	hystrix.Configure([]string{"laracom.service.demo.DemoService.SayHello"})
	service := micro.NewService(
		micro.Name("laracom.demo.cli"),
		micro.WrapClient(hystrix.NewClientWrapper()),
	)
	service.Init()

	client := pb.NewDemoServiceClient("laracom.service.demo", service.Client())

	// 模拟常驻内存
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case os.Interrupt, os.Kill, syscall.SIGQUIT:
				fmt.Println("退出客户端")
				os.Exit(0)
			default:
				fmt.Println("程序执行中...")
			}
		}
	}()

	for {
		rsp, err := client.SayHello(context.TODO(), &pb.DemoRequest{Name: "学院君"})
		if err != nil {
			log.Fatalf("服务调用失败：%v", err)
			return
		}
		log.Println(rsp.Text)
		time.Sleep(3 * time.Second)
	}
}