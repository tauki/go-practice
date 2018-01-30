package main

import (
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
	test.Post("/").
		SetHeader("Content-Type", "application/json").
			Expect(t).
				Type("json").
					JSONSchema(Schema).
						Done()
}