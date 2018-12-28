package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "tauki.com/practice/merger-redis/protobuffer"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Merge(ctx context.Context, in *pb.Data) (*pb.Data, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "pass",
		DB:       0, // use default DB
	})

	val, err := client.Get("old").Result()
	if err != nil {
		panic(err)
	}

	var oldData pb.Data
	err = json.Unmarshal([]byte(val), &oldData)
	if err != nil {
		panic(err)
	}

	for _, entity := range in.Entity {
		check := true
		for _, old := range oldData.Entity {
			if entity.Code == old.Code {
				check = false
				break
			}
		}
		if check == true {
			oldData.Entity = append(oldData.Entity, entity)
		}
	}

	newData, err := json.Marshal(oldData)
	err = client.Set("old", newData, 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successful!")

	return &oldData, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMergerServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
