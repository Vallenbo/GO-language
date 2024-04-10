package main

import (
	"GoNotebook/Go_grpc/grpc四种模式/双向端流服务模式/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) SquareUpdates(stream proto.Calculator_SquareUpdatesServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		result := req.Number * req.Number
		if err := stream.Send(&proto.SquareResponse{Result: result}); err != nil {
			return err
		}
	}
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
