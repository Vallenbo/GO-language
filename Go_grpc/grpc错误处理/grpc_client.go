package main

import (
	"GoNotebook/Go_grpc/hello_grpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。 此处使用不安全的证书来实现 SSL/TLS 连接
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr :8080 连接失败 %s", err))
	}
	defer conn.Close()

	client := hello_grpc.NewHelloServiceClient(conn) // 初始化客户端
	ResponseBody, err := client.SayHello(context.Background(), &hello_grpc.HelloRequest{
		Name:    "枫枫",
		Message: "ok",
	})
	if err != nil {
		st, _ := status.FromError(err)
		fmt.Println("Code", st.Code(), "Message") //  提取返回err中的状态码和信息
	}
	fmt.Println(ResponseBody.GetName())
}
