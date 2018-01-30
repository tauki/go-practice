package main

import (
	"github.com/go-redis/redis"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	pb "tauki.com/practice/merger-redis/protobuffer"
)

func readData() pb.Data {
	var data pb.Data

	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	oldData, _ := ioutil.ReadFile("old_data.json")

	json.Unmarshal(oldData, &data)
	return data
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "pass", // no password set
		DB:       0,  // use default DB
	})

	data := readData()
	s, err := json.Marshal(data)
	client.Del("old")
	err = client.Set("old", s, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successful!")
}
