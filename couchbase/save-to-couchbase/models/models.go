package models

type Config struct {
	// Server Information

	Environment string
	ServePort   string
	ServeHost   string

	// couchbase

	BucketName string
	BucketPass string
	DBHost     string
	DBUser     string
	DBPass     string
	DBPort     string
}
