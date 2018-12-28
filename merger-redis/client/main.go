package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	pb "github.com/tauki/go-practice/merger/protobuffer"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
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

func (h *handler) get(gc *gin.Context) {
	c := pb.NewMergerClient(h.client)

	data := readData()
	r, err := c.Merge(context.Background(), &data)
	if err != nil {
		gc.JSON(http.StatusNoContent, err)
		return
	}
	gc.JSON(http.StatusOK, r)
}

func (h *handler) post(gc *gin.Context) {
	c := pb.NewMergerClient(h.client)
	var data *pb.Data
	gc.BindJSON(&data)
	r, err := c.Merge(context.Background(), data)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	gc.JSON(http.StatusOK, r)
}

func readData() pb.Data {
	var data pb.Data

	//pwd, _ := os.Getwd()
	//fmt.Println(pwd)

	oldData, _ := ioutil.ReadFile("new_data.json")

	json.Unmarshal(oldData, &data)
	return data
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	defer conn.Close()

	h := getHandler(conn)

	//http.HandleFunc("/", h.send)
	//http.ListenAndServe(":9000", nil)

	router := gin.Default()

	router.GET("/", h.get)
	router.POST("/", h.post)

	router.Run(":9000")
}
