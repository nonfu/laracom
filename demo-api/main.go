package main

import (
    "context"
    "encoding/json"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/errors"
    "log"
    "strings"
    pb "github.com/nonfu/laracom/demo-service/proto/demo"
    api "github.com/micro/go-micro/api/proto"
)

type Greeter struct {
    Client pb.DemoServiceClient
}

func (s *Greeter) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
    log.Println("收到 /greeter/hello API 请求")

    // 从请求参数中获取 name 值
    name, ok := req.Get["name"]
    if !ok || len(name.Values) == 0 {
        return errors.BadRequest("laracom.api.demo", "名字不能为空")
    }

    // 将参数交由底层服务处理
    response, err := s.Client.SayHello(ctx, &pb.DemoRequest{
        Name: strings.Join(name.Values, " "),
    })
    if err != nil {
        return err
    }

    // 处理成功，则返回处理结果
    rsp.StatusCode = 200
    b, _ := json.Marshal(map[string]string{
        "message": response.Text,
    })
    rsp.Body = string(b)

    return nil
}

func main() {
    // 创建一个新的服务
    service := micro.NewService(
        micro.Name("laracom.api.demo"),
    )
    service.Init()

    // 将请求转发给底层 laracom.service.demo 服务处理
    service.Server().Handle(
        service.Server().NewHandler(
            &Greeter{Client: pb.NewDemoServiceClient("laracom.service.demo", service.Client())},
        ),
    )

    // 运行服务
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
