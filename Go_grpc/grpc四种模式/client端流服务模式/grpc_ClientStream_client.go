package main

import (
	"GoNotebook/Go_grpc/grpc四种模式/client端流服务模式/proto"
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewCalculatorClient(conn)

	stream, err := client.SquareBatch(context.Background())
	if err != nil {
		log.Fatalf("error while calling SquareBatch RPC: %v", err)
	}

	numbers := []int32{2, 3, 4, 5}
	for _, number := range numbers {
		if err := stream.Send(&proto.SquareRequest{Number: number}); err != nil {
			log.Fatalf("error while sending request: %v", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %v", err)
	}
	log.Printf("Result of batch square is: %d", response.Result)
}
