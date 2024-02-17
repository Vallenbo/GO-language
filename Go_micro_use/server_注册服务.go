package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
)

type ExampleService struct{}

func (e *ExampleService) Call(ctx context.Context, req *string, rsp *string) error {
	*rsp = "Hello, " + *req
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("example.service"),
	)

	service.Init()

	micro.RegisterHandler(service.Server(), new(ExampleService))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
