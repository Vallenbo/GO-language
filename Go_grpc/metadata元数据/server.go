package main

import (
	"GoNotebook/Go_grpc/hello_grpc"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, req *hello_grpc.HelloRequest) (*hello_grpc.HelloResponse, error) {
	//获取header
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("get metadata error")
	}
	for key, val := range md {
		fmt.Println(key, val)
	}
	//获取header中的name*********
	//if nameSlice, ok := md["name"]; ok {
	//	fmt.Println(nameSlice)
	//	for i, e := range nameSlice {
	//		fmt.Println(i, e)
	//	}
	//}
	return &hello_grpc.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	hello_grpc.RegisterHelloServiceServer(g, &Server{})
	lis, err := net.Listen("tcp", "127.0.0.1:8083")
	if err != nil {
		panic("failed to listen：" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc：" + err.Error())
	}
}
