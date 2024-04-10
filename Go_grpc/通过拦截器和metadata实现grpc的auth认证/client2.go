package main

import (
	"GoNotebook/Go_grpc/hello_grpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ClientCredentials struct{}

func (c ClientCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":    "111111",
		"password": "123456",
	}, nil
}

func (c ClientCredentials) RequireTransportSecurity() bool {
	return false
}

func main() {
	addr := ":8080" // gRPC服务器地址
	// 创建客户端拦截器，拦截器在每次gRPC调用之前运行，用于记录调用的耗时
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(ClientCredentials{})) // 创建带有拦截器的gRPC选项

	// 通过地址和选项创建gRPC连接
	// 使用grpc.WithInsecure()选项表示不使用TLS
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server at [%s]: %s", addr, err) // 如果连接失败，打印错误并退出
	}
	defer conn.Close() // 确保在函数退出时关闭连接

	// 使用连接初始化gRPC客户端
	client := hello_grpc.NewHelloServiceClient(conn)

	// 调用gRPC服务的SayHello方法
	// 使用context.Background()作为上下文，表示该调用没有超时和取消信号
	// 使用&hello_grpc.HelloRequest{Name: "枫枫", Message: "ok"}作为请求参数
	result, err := client.SayHello(context.Background(), &hello_grpc.HelloRequest{
		Name:    "枫枫",
		Message: "ok",
	})

	// 打印调用结果和可能的错误
	fmt.Println(result, err)
}
