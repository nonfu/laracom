package main

import (
    pb "github.com/nonfu/laracom/demo-service/proto/demo"
    . "github.com/smartystreets/goconvey/convey"
    "golang.org/x/net/context"
    "testing"
)

func TestDemoServiceHandler_SayHello(t *testing.T) {
    Convey("Given a test for DemoService.SayHello", t, func() {
        context := context.Background();
        service := &DemoServiceHandler{}

        Convey("Get greeting text with specified name", func() {
            var resp pb.DemoResponse
            service.SayHello(context, &pb.DemoRequest{Name: "学院君"}, &resp)

            Convey("Then the response text should be a '你好, 学院君'", func() {
                So(resp.Text, ShouldEqual, "你好, 学院君")
            })
        })

        Convey("Get greeting text without name", func() {
            var resp pb.DemoResponse
            service.SayHello(context, &pb.DemoRequest{}, &resp)

            Convey("Then the response text should be '你好, '", func() {
                So(resp.Text, ShouldEqual, "你好, ")
            })
        })
    })
}