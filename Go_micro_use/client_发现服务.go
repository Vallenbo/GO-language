package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
)

func main() {
	service := micro.NewService(micro.Name("example.client"))
	service.Init()

	exampleService := micro.Client(stNewClient().Service("example.service"))

	req := "World"
	var rsp string

	err := exampleService.Call(context.Background(), &req, &rsp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rsp)
	}
}
