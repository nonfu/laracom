package handler

import (
    "github.com/nonfu/laracom/user-service/model"
    pb "github.com/nonfu/laracom/user-service/proto/user"
    "github.com/nonfu/laracom/user-service/repo/mocks"
    . "github.com/smartystreets/goconvey/convey"
    "golang.org/x/net/context"
    "testing"
)

func TestUserService_Get(t *testing.T) {
    Convey("Given a test for DemoService.SayHello", t, func() {
        ctx := context.Background()

        // 初始化 Repository 模拟类实例
        mockRepo := &mocks.Repository{}
        // 初始化模拟接口方法返回值
        mockRepo.On("Get", uint(1)).Return(&model.User{Name:"test"}, nil)
        mockRepo.On("GetByEmail", "test@xueyuanjun.com").Return(&model.User{Email: "test@xueyuanjun.com"}, nil)

        service := &UserService{Repo: mockRepo}

        Convey("When get user with specified id", func() {
            var resp pb.Response
            service.Get(ctx, &pb.User{Id: "1"}, &resp)

            Convey("Then the response user name should be test", func() {
                So(resp.User.Name, ShouldEqual, "test")
            })
        })

        Convey("When get user with specified email", func() {
            var resp pb.Response
            service.Get(ctx, &pb.User{Email: "test@xueyuanjun.com"}, &resp)

            Convey("Then the response user email should be  'test@xueyuanjun.com", func() {
                So(resp.User.Email, ShouldEqual, "test@xueyuanjun.com")
            })
        })
    })
}
