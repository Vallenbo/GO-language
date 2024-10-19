package main

import (
	"GoNotebook/Go_grpc/hello_grpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	"net"
)

// HelloServer1 得有一个结构体，需要实现这个服务的全部方法,叫什么名字不重要
type HelloServer struct {
}

func (*HelloServer) SayHello(ctx context.Context, request *hello_grpc.HelloRequest) (*hello_grpc.HelloResponse, error) {
	fmt.Println("入参：", request.Name, request.Message)
	//return nil, status.Error(codes.NotFound, "记录未找到")
	return nil, status.Errorf(codes.NotFound, "记录未找到:%s", request.Name)
}

func main() {
	listen, err := net.Listen("tcp", ":8080") // 监听端口
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer() // 创建一个gRPC服务器实例。
	server := HelloServer{}
	hello_grpc.RegisterHelloServiceServer(s, &server) // 将server结构体注册为gRPC服务。
	fmt.Println("grpc server running :8080")

	err = s.Serve(listen) // 开始处理客户端请求。
}
