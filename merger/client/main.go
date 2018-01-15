package main


import (
	pb "tauki.com/practice/merger/protobuffer"
	"google.golang.org/grpc"
	"log"
	"context"
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
	"net/http"
)

const (
	address = "localhost:50051"
)

type handler struct {
	client *grpc.ClientConn
}

func getHandler(c *grpc.ClientConn) *handler {
	return &handler{client: c}
}

func (h *handler) send(w http.ResponseWriter, req *http.Request) {
	c := pb.NewMergerClient(h.client)

	data := readData()
	r, err := c.Merge(context.Background(), &data)
	if err != nil {
		log.Fatalf("could not calculate %v", err)
	}

	log.Println(r)
	json.NewEncoder(w).Encode(r)
}

func readData() (pb.Data) {
	var data pb.Data

	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	oldData, _ := ioutil.ReadFile("new_data.json")

	json.Unmarshal(oldData, &data)
	return data
}

func main () {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	defer conn.Close()

	h := getHandler(conn)

	http.HandleFunc("/", h.send)

	http.ListenAndServe(":9000", nil)
}