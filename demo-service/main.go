package main

import (
	"context"
	ratelimit "github.com/juju/ratelimit"
	"github.com/micro/go-micro"
	ratelimiter "github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
	pb "github.com/nonfu/laracom/demo-service/proto/demo"
	userpb "github.com/nonfu/laracom/user-service/proto/user"
	"log"
)

const QPS = 1000

type DemoServiceHandler struct {

}

func (s *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	rsp.Text = "你好, " + req.Name
	return nil
}

func (s *DemoServiceHandler) SayHelloByUserId(ctx context.Context, req *pb.HelloRequest, rsp *pb.DemoResponse) error {
	service := micro.NewService()
	service.Init()
	client := userpb.NewUserServiceClient("laracom.service.user", service.Client())
	resp, err := client.Get(context.TODO(), &userpb.User{Id: req.Id})
	if err != nil {
		return err
	}
	rsp.Text = "你好, " + resp.User.Name
	return nil
}

func main()  {
	bucket := ratelimit.NewBucketWithRate(float64(QPS), int64(QPS))

	// 注册服务名必须和 demo.proto 中的 package 声明一致
	service := micro.NewService(
		micro.Name("laracom.service.demo"),
		micro.WrapHandler(ratelimiter.NewHandlerWrapper(bucket, false)),
	)
	service.Init()

	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}