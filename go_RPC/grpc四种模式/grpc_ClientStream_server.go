package main

import (
	"GoNotebook/go_RPC/grpc四种模式/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ClientStream struct{}

func (ClientStream) UploadFile(stream proto.ClientStream_UploadFileServer) error {
	for i := 0; i < 10; i++ {
		response, err := stream.Recv()
		fmt.Println(response, err)
	}
	stream.SendAndClose(&proto.Response{Text: "完毕了"})
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	proto.RegisterClientStreamServer(server, &ClientStream{})

	server.Serve(listen)
}
