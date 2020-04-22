package main

import (
    "context"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/web"
    pb "github.com/nonfu/laracom/demo-service/proto/demo"
)

type Say struct{}

var (
    cli pb.DemoServiceClient
)

func (s *Say) Anything(c *gin.Context) {
    log.Print("Received Say.Anything API request")
    c.JSON(200, map[string]string{
        "message": "你好，学院君",
    })
}

func (s *Say) Hello(c *gin.Context) {
    log.Print("Received Say.Hello API request")

    name := c.Param("name")

    response, err := cli.SayHello(context.TODO(), &pb.DemoRequest{
        Name: name,
    })

    if err != nil {
        c.JSON(500, err)
    }

    c.JSON(200, response)
}

func main() {
    // Create service
    service := web.NewService(
        web.Name("go.micro.api.greeter"),
    )

    service.Init()

    // setup Greeter Server Client
    cli = pb.NewDemoServiceClient("laracom.service.demo", client.DefaultClient)

    // Create RESTful handler (using Gin)
    say := new(Say)
    router := gin.Default()
    router.GET("/greeter", say.Anything)
    router.GET("/greeter/:name", say.Hello)

    // Register Handler
    service.Handle("/", router)

    // Run server
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}