package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "tauki.com/practice/hello-grpc/protobuffer"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) Add(ctx context.Context, in *pb.Calc) (*pb.Sum, error) {
	return &pb.Sum{Sum: in.F + in.S}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
