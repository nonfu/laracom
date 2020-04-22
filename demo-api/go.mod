module github.com/nonfu/laracom/demo-api

go 1.13

replace github.com/nonfu/laracom/common v0.0.0-20200422074139-c68e2b3d4434 => /Users/sunqiang/Development/go/src/laracom/common

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/micro/go-micro v1.18.0
	github.com/nonfu/laracom/demo-service v0.0.0-20200420164645-fee8b63eddb6
)
