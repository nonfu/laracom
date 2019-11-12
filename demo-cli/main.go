package main

import (
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("laracom.demo.cli"))
	service.Init()

}