package main

import (
	"context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	pb "tauki.com/practice/filesend/protobuffer"
)

const (
	port      = ":50051"
	directory = "test/"
)

type server struct{}

func (s *server) SendFile(ctx context.Context, in *pb.Files) (*pb.Resp, error) {
	resp := pb.Resp{Resp: "successful"}

	for _, file := range in.Files {
		err := ioutil.WriteFile(directory+file.FileName, file.File, 0777)
		if err != nil {
			log.Fatal(err)
			return &pb.Resp{Resp: "failed"}, err
		}
	}

	return &resp, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFileSenderServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
