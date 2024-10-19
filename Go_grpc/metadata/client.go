package main

import (
	"GoNotebook/Go_grpc/hello_grpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial(":8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := hello_grpc.NewHelloServiceClient(conn)

	//写入metadata
	md := metadata.New(map[string]string{
		"name":     "lff",
		"password": "123456",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.SayHello(ctx, &hello_grpc.HelloRequest{Name: "lff111"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
