package main

import (
	"GoNotebook/Go_grpc/hello_grpc"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
)

func main() {
	msg := &hello_grpc.HelloRequest{
		Name:    "123",
		Message: "123123",
	}

	v, err := protovalidate.New()
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	if err = v.Validate(msg); err != nil {
		fmt.Println("validation failed:", err)
	} else {
		fmt.Println("validation succeeded")
	}
}
