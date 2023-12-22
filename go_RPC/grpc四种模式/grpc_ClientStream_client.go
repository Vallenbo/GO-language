package main

import (
	"GoNotebook/go_RPC/grpc四种模式/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"
)

func main() {
	addr := ":8080"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	defer conn.Close()
	// 初始化客户端
	client := proto.NewClientStreamClient(conn)
	stream, err := client.UploadFile(context.Background())
	for i := 0; i < 10; i++ {
		stream.Send(&proto.FileRequest{FileName: fmt.Sprintf("第%d次", i)})
	}
	response, err := stream.CloseAndRecv()
	fmt.Println(response, err)
}
