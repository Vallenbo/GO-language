package main

import (
	"GoNotebook/Go_grpc/grpcFourMode/ServerAndClientStreamMode/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewCalculatorClient(conn)

	stream, err := client.SquareUpdates(context.Background())
	if err != nil {
		log.Fatalf("error while calling SquareUpdates RPC: %v", err)
	}

	numbers := []int32{2, 3, 4, 5}
	for _, number := range numbers {
		if err := stream.Send(&proto.SquareRequest{Number: number}); err != nil {
			log.Fatalf("error while sending request: %v", err)
		}
		time.Sleep(time.Second) // Simulate processing time
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("error while receiving response: %v", err)
		}
		log.Printf("Square of %d is: %d", number, response.Result)
	}
	stream.CloseSend()
}
