package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"GoNotebook/Go_grpc/grpcFourMode/serverStreamMode/proto"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewCalculatorClient(conn)

	number := int32(5)
	stream, err := client.SquareStream(context.Background(), &proto.SquareRequest{Number: number})
	if err != nil {
		log.Fatalf("error while calling SquareStream RPC: %v", err)
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("error while receiving response: %v", err)
		}
		log.Printf("Square of %d is: %d", response.Result, response.Result)
	}
}
