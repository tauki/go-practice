package main

import (
	"fmt"
	"github.com/urfave/cli"
	"tauki.com/practice/couchbase/save-to-couchbase/models"
	"tauki.com/practice/couchbase/save-to-couchbase/router"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var cfg *models.Config

func init() {
	env := getEnv("ENV", "production")
	dbuser, dbpass := getCred()

	switch env {

	default:
		cfg = &models.Config{
			Environment: env,
			ServePort:   "9090",
			ServeHost:   "0.0.0.0",

			BucketName: "user",
			BucketPass: "",
			DBPort:     "8091",
			DBUser:     dbuser,
			DBPass:     dbpass,
			DBHost:     "localhost",
			// ...
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "ScraperDB"
	app.Usage = "Database endpoints for storing and accessing scraped data"
	app.Version = fmt.Sprintf("%s-alpha", time.Now().Local().Format("00.00.00"))
	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Run ScraperDB app",
			Action: func(c *cli.Context) {
				route, err := router.InitRouter(cfg)
				if err != nil {
					errorMsg(err, "Router")
				} else {
					err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.ServeHost, cfg.ServePort), route)
					if err != nil {
						errorMsg(err, "Serve")
					}
				}
			},
		},
	}
	app.Run(os.Args)
}

func errorMsg(err error, context string) {
	msg := fmt.Sprintf("main :: %s :: %s", context, err)
	panic(msg)
}

func getEnv(key, defaults string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaults
}

func getCred() (string, string) {
	f, err := ioutil.ReadFile("keys/dbcreds")
	if err != nil {
		errorMsg(err, "getCred")
	}
	creds := strings.Split(string(f), " ")
	if len(creds) < 2 {
		panic("check credentials")
	}

	// suffix of the last string may contain an extra
	return creds[0], strings.TrimSuffix(creds[1], "\n")
}
