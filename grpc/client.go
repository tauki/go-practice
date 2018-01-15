package main

import (
	pb "tauki.com/practice/grpc/proto"
	"log"
	"github.com/golang/protobuf/proto"
	"fmt"
)



func main () {
	comp := &pb.Company{
		Name: "Example Corp",
		Adress : &pb.AddressInfo{
			City:    "London",
			Country: "UK",
		},
		Type: pb.CompanyType_Private,
		Employees: []*pb.Employee{
			&pb.Employee{
				Name: "John",
			},
		},
	}

	data, err := proto.Marshal(comp)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	serialized := &pb.Company{}
	err = proto.Unmarshal(data, serialized)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println(data)
	fmt.Println(serialized)
}