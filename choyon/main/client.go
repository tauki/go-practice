package main

import (
	"context"
	"google.golang.org/grpc"
	pb "ideeza/emergebox/buffers/auth"
	"log"
)

func main() {
	conn, err := grpc.Dial("ADDRESS", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	defer conn.Close()

	c := pb.NewAuthClient(conn)

	u, err := c.UserSignup(context.Background(), &pb.UserSignupRequest{
		UserInfo: &pb.UserInfo{
			User: &pb.User{
				FirstName: "Abdul",
				LastName:  "Kudzus",
				Password:  "amikuddus",
				Email:     "abdulkuddus@user.com",
			},
		},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(u)
	err = nil

	sp, err := c.SPSignup(context.Background(), &pb.ServiceProviderSignupRequest{
		SpInfo: &pb.ServiceProviderInfo{
			User: &pb.User{
				FirstName: "Abdul",
				LastName:  "Kudzus",
				Password:  "amikuddus",
				Email:     "abdul@kuddus.com",
			},
			PhoneNumber: "00112233449955",
			CompanyName: "Kuddus Co",
			City:        "dhaka'r baaire",
		},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(sp)

	l, err := c.UserLogin(context.Background(), &pb.UserLoginRequest{
		User: &pb.User{
			FirstName: "Abdul",
			LastName:  "Kudzus",
			Password:  "amikuddus",
			Email:     "abdulkuddus@user.com",
		},
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println(l)
	}

	spl, err := c.SPLogin(context.Background(), &pb.ServiceProviderLoginRequest{
		User: &pb.User{
			FirstName: "Abdul",
			LastName:  "Kudzus",
			Password:  "amikuddus",
			Email:     "abdul@kuddus.com",
		},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(spl)
}
