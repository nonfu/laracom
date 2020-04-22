package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/json"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/metadata"
	traceplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/nonfu/laracom/common/wrapper/breaker/hystrix"
	pb "github.com/nonfu/laracom/demo-service/proto/demo"
	"github.com/nonfu/laracom/common/tracer"
	userpb "github.com/nonfu/laracom/user-service/proto/user"
	"github.com/opentracing/opentracing-go"
	"log"
	"os"
	"strings"
	"time"
)

type DemoServiceHandler struct {
	appConfig *AppConfig
}

func (s *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	// 从微服务上下文中获取追踪信息
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// 创建新的 Span 并将其绑定到微服务上下文
	sp = opentracing.StartSpan("SayHello", opentracing.ChildOf(wireContext))
	// 记录请求
	sp.SetTag("req", req)
	defer func() {
		// 记录响应
		sp.SetTag("res", rsp)
		// 在函数返回 stop span 之前，统计函数执行时间
		sp.Finish()
	}()

	// 如果参数为空设置默认值
	if req.Name == "" {
		req.Name = s.appConfig.UserName
	}
	rsp.Text = "你好, " + req.Name
	return nil
}

func (s *DemoServiceHandler) SayHelloByUserId(ctx context.Context, req *pb.HelloRequest, rsp *pb.DemoResponse) error {

	// 使用断路器
	hystrix.Configure([]string{"laracom.service.user.UserService.GetById"})
	service := micro.NewService(
		micro.WrapClient(hystrix.NewClientWrapper()),
	)
	client := userpb.NewUserServiceClient("laracom.service.user", service.Client())
	resp, err := client.GetById(context.TODO(), &userpb.User{Id: req.Id})
	if err != nil {
		return err
	}
	rsp.Text = "你好, " + resp.User.Name
	return nil
}

func main()  {
	// 获取viper配置实例
	appConfig := initAppConfig()

	// 初始化全局服务追踪
	t, io, err := tracer.NewTracer(appConfig.ServiceName, os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 注册服务名必须和 demo.proto 中的 package 声明一致
	service := micro.NewService(
		micro.Name(appConfig.ServiceName),
		micro.WrapHandler(traceplugin.NewHandlerWrapper(opentracing.GlobalTracer())), // 基于 jaeger 采集追踪数据
	)
	service.Init()

	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{appConfig: appConfig})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}

func initAppConfig() *AppConfig {
	// 默认使用 JSON 编码
	encoder := json.NewEncoder()
	fileSource := file.NewSource(file.WithPath("./demo.json"), source.WithEncoder(encoder))
	etcdSource := etcd.NewSource(
		etcd.WithAddress(strings.Split(os.Getenv("MICRO_REGISTRY_ADDRESS"), ",")[0]),
		etcd.WithPrefix("laracom/demo/"),
		source.WithEncoder(encoder),
	)
	conf := config.NewConfig()
	var err error
	if os.Getenv("ENABLE_REMOTE_CONFIG") == "true" {
		err = conf.Load(
			fileSource,  // 将文件配置作为默认值
			etcdSource,  // 会覆盖上面的文件配置
		)
	} else {
		err = conf.Load(fileSource)
	}
	if err != nil {
		// 加载数据源失败
		log.Fatalf("读取配置失败: %v", err)
	}
	var appConfig AppConfig
	err = conf.Get("laracom", "demo", "app").Scan(&appConfig)
	if err != nil {
		// 读取远程配置失败
		log.Fatalf("读取配置失败: %v", err)
	}
	log.Printf("初始化配置：%v", appConfig)
	log.Printf("初始化配置：%v", conf.Map())

	// 开启协程监听配置变更
	go func(){
		for {
			time.Sleep(time.Second * 5) // delay after each request

			w, err := conf.Watch("laracom", "demo", "app")
			if err != nil {
				log.Printf("监听配置变更失败: %v", err)
				continue
			}

			// wait for next value
			value, err := w.Next()
			if err != nil {
				log.Printf("读取配置变更失败: %v", err)
				continue
			}

			value.Scan(&appConfig)
			log.Printf("配置值变更：%s", &appConfig)
		}
	}()

	return &appConfig
}

type DemoConfig struct {
	App AppConfig `json:"app"`
}

type AppConfig struct {
	ServiceName string `json:"service_name"`
	UserName string `json:"user_name"`
}