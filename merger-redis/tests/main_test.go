package main

import (
	pb "github.com/tauki/go-practice/merger-redis/protobuffer"
	"gopkg.in/h2non/baloo.v3"
	"testing"
)

var test = baloo.New("http://localhost:9000")

func TestMergerGetResponse(t *testing.T) {
	test.Get("/").
		SetHeader("Content-Type", "Bar").
		Expect(t).
		Type("json").
		JSONSchema(Schema).
		Done()
}

func TestMergerPostResponse(t *testing.T) {
	var data pb.Data

	data.Entity = append(data.Entity, &pb.Entity{
		Code:    1,
		REF:     "color",
		Display: "Runway Hit",
	})

	test.Post("/").
		SetHeader("Content-Type", "application/json").
		JSON(data).
		Expect(t).
		Type("json").
		JSONSchema(Schema).
		Done()
}
