package main

import (
	"context"
	"fmt"
	pb "github.com/tauki/go-practice/filesend/protobuffer"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
)

const (
	address   = "localhost:50051"
	directory = "test/"
)

func getFiles() *pb.Files {
	dir, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	var files []*pb.File
	for _, fileInfo := range dir {
		f, err := ioutil.ReadFile(directory + fileInfo.Name())
		if err != nil {
			log.Fatal(err)
		}

		files = append(files, &pb.File{
			File:     f,
			FileName: fileInfo.Name(),
		})
	}

	return &pb.Files{
		Files: files,
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	defer conn.Close()

	c := pb.NewFileSenderClient(conn)
	r, err := c.SendFile(context.Background(), getFiles())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
