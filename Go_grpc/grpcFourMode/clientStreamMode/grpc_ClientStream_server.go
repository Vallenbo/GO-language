package main

import (
	"GoNotebook/Go_grpc/grpcFourMode/clientStreamMode/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) SquareBatch(stream proto.Calculator_SquareBatchServer) error {
	var result int32 = 1
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		result *= req.Number
	}
	return stream.SendAndClose(&proto.SquareResponse{Result: result})
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
