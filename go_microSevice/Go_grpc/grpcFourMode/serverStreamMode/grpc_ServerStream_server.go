package main

import (
	"GoNotebook/Go_grpc/grpcFourMode/serverStreamMode/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct{}

func (s *server) SquareStream(req *proto.SquareRequest, stream proto.Calculator_SquareStreamServer) error {
	for i := 1; i < int(req.Number); i++ {
		result := i * i
		if err := stream.Send(&proto.SquareResponse{Result: int32(result)}); err != nil {
			return err
		}
		time.Sleep(time.Second) // Simulate processing time
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterCalculatorServer(s, &server{})
	log.Println("Server started at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
