# product-service/Dockerfile

# 使用 golang 官方镜像提供 Go 运行环境，并且命名为 buidler 以便后续引用
FROM golang:alpine as builder

# 启用 Go Modules 并设置 GOPROXY
ENV GO111MODULE on

# 安装 git
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
	apk update && \
	apk add --no-cache git

# 设置工作目录
WORKDIR /app/laracom-product-service

# 将当前项目所在目录代码拷贝到镜像中
COPY . .

# 下载依赖
RUN go env -w GOPROXY=https://goproxy.io,direct && go mod tidy

# 构建二进制文件，添加来一些额外参数以便可以在 Alpine 中运行它
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o laracom-product-service

# 下面是第二阶段的镜像构建，和之前保持一致
FROM alpine:latest

# 更新依赖软件
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
	apk update && \
	apk add --no-cache bash ca-certificates

# 和上个阶段一样设置工作目录
RUN mkdir /app

WORKDIR /app

# 这一步不再从宿主机拷贝二进制文件，而是从上一个阶段构建的 builder 容器中拉取，神奇吧
COPY --from=builder /app/laracom-product-service/laracom-product-service .

# 启动商品服务
CMD ["./laracom-product-service"]