package main

import (
	"GoNotebook/go_RPC/hello_grpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	addr := ":8080"
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	defer conn.Close()
	// 初始化客户端
	client := hello_grpc.NewHelloServiceClient(conn)
	ctxDeadline, _ := context.WithTimeout(context.Background(), time.Second*3)
	result, err := client.SayHello(ctxDeadline, &hello_grpc.HelloRequest{
		Name:    "枫枫",
		Message: "ok",
	})
	fmt.Println(result, err)
}
