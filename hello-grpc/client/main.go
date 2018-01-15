package main


import (
	pb "tauki.com/practice/hello-grpc/protobuffer"
	"google.golang.org/grpc"
	"log"
	"context"
)

const (
	address = "localhost:50051"
)

func main () {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	r, err := c.Add(context.Background(), &pb.Calc{F: 1, S: 2})
	if err != nil {
		log.Fatalf("could not calculate %v", err)
	}
	log.Printf("Sum : %d", r.Sum)
}