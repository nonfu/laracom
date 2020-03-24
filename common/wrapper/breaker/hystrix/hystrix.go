package hystrix

import (
    "github.com/afex/hystrix-go/hystrix"
    "github.com/micro/go-micro/client"
    "log"
    "net"
    "net/http"

    "context"
)

type clientWrapper struct {
    client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
    return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
        return c.Client.Call(ctx, req, rsp, opts...)
    }, nil)
}

// NewClientWrapper returns a hystrix client Wrapper.
func NewClientWrapper() client.Wrapper {
    return func(c client.Client) client.Client {
        return &clientWrapper{c}
    }
}

func Configure(names []string) {
    // Hystrix 有默认的参数配置，这里可以针对某些 API 进行自定义配置
    config := hystrix.CommandConfig{
        Timeout:               3000,
        MaxConcurrentRequests: 100,
        ErrorPercentThreshold: 25,
    }
    configs := make(map[string]hystrix.CommandConfig)
    for _, name := range names {
        configs[name] = config
    }
    hystrix.Configure(configs)

    // 结合 Hystrix Dashboard 将服务状态信息可视化
    hystrixStreamHandler := hystrix.NewStreamHandler()
    hystrixStreamHandler.Start()
    go http.ListenAndServe(net.JoinHostPort("", "8181"), hystrixStreamHandler)
}
