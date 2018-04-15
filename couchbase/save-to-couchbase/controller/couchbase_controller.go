package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/couchbase/gocb"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"tauki.com/practice/couchbase/save-to-couchbase/models"
	"log"
	"net/http"
	"strings"
)

type BucketServer struct {
	Document *gocb.Bucket
	cfg      *models.Config
}

func GetCouchDBController(cfg *models.Config) (*BucketServer, error) {
	cluster, err := gocb.Connect(cfg.DBHost + ":" + cfg.DBPort)
	if err != nil {
		fmt.Println(err)
		return nil, errorMessage(err, "GetController")
	}
	err = cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: cfg.DBUser,
		Password: cfg.DBPass,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	bucket, err := cluster.OpenBucket(cfg.BucketName, cfg.BucketPass)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &BucketServer{
		Document: bucket,
		cfg:      cfg,
	}, nil
}

func (b BucketServer) GetAll(c *gin.Context) {
	query := fmt.Sprintf("SELECT * FROM `%s` as couch;", b.cfg.BucketName)
	rows, err := b.Document.ExecuteN1qlQuery(gocb.NewN1qlQuery(query), nil)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	var row map[string]interface{}
	var data []map[string]interface{}

	for rows.Next(&row) {
		data = append(data, row)
		// need to make row empty or it will just keep re-writing on the same memory while /
		// adding the next element to the same memory so while the added memory to data being /
		// a pointer, is pointing to the same memory address making it appear as the same element /
		// while it's actually just pointing to the same memory address
		row = nil //todo: try to find a better way to fix this
	}

	if len(data) == 0 {
		c.JSON(http.StatusNotFound,
			http.StatusText(http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, data)
}

// todo: no unique constraint so 2 same values are currently stored
func (b BucketServer) Create(c *gin.Context) {
	var n1qlparams []interface{}
	var data interface{}

	//todo: fix this
	_ = json.NewDecoder(c.Request.Body).Decode(&data)
	JSON, _ := json.Marshal(data)

	id, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		log.Println(err)
	}

	n1qlparams = append(n1qlparams, id)

	query := "INSERT INTO `" + b.cfg.BucketName + "` (KEY, VALUE) VALUES ($1, " + string(JSON) + ")"
	_, err = b.Document.ExecuteN1qlQuery(gocb.NewN1qlQuery(query), n1qlparams)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusCreated, http.StatusText(http.StatusCreated))
}

func (b BucketServer) GetByID(c *gin.Context) {
	var data interface{}

	id := c.Param("id")

	query := fmt.Sprintf("SELECT * FROM `%s` as couch WHERE id=%s;", b.cfg.BucketName, id)
	rows, err := b.Document.ExecuteN1qlQuery(gocb.NewN1qlQuery(query), nil)
	if err != nil {
		c.JSON(http.StatusNotFound,
			http.StatusText(http.StatusNotFound))
		return
	}
	rows.One(&data) // ID is suppose to be unique though not implemented yet. todo
	if data == nil {
		c.JSON(http.StatusNotFound,
			http.StatusText(http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, data)
}

/*
select * from `test-scraperdb` WHERE
ARRAY_CONTAINS(tags, "car") = true AND NOT ARRAY_CONTAINS(tags, "test") = true;
*/
func (b BucketServer) Search(c *gin.Context) {
	params := c.Request.URL.Query()
	if len(params) == 0 {
		c.JSON(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest))
		return
	}

	query := fmt.Sprintf("SELECT * FROM `%s` as couch", b.cfg.BucketName)

	if len(params["name"]) != 0 {
		query += " " + fmt.Sprintf("WHERE lower(name) LIKE %s", "\"%"+params["name"][0]+"%\"")
	}

	if len(params["tag"]) != 0 {

		if len(params["name"]) != 0 {
			query += " AND"
		} else {
			query += " WHERE"
		}

		tags := strings.Split(params["tag"][0], ",")
		for i, v := range tags {
			query = query + " " + fmt.Sprintf("ARRAY_CONTAINS(tags, \"%s\") = true", string(v))
			if i < len(tags)-1 {
				query += " AND"
			}
		}
	}

	if len(params["limit"]) != 0 {
		query += " " + fmt.Sprintf("LIMIT %s", params["limit"][0])
	}

	rows, err := b.Document.ExecuteN1qlQuery(gocb.NewN1qlQuery(query), nil)
	if err != nil {
		c.JSON(http.StatusNotFound,
			http.StatusText(http.StatusNotFound))
		return
	}

	var row map[string]interface{}
	var data []map[string]interface{}

	for rows.Next(&row) {
		data = append(data, row)
		// need to make row empty or it will just keep re-writing on the same memory while /
		// adding the next element to the same memory so while the added memory to data being /
		// a pointer, is pointing to the same memory address making it appear as the same element /
		// while it's actually just pointing to the same memory address
		row = nil //todo: try to find a better way to fix this
	}

	if len(data) == 0 {
		c.JSON(http.StatusNotFound,
			http.StatusText(http.StatusNotFound))
		return
	}
	c.JSON(http.StatusOK, data)
}

/*
Query to couchbase using post-request
*/

type Request struct {
	Query string `json:"query"`
}

func (b BucketServer) Query(c *gin.Context) {
	var req Request

	if err := c.ShouldBindJSON(&req); err == nil {
		rows, err := b.Document.ExecuteN1qlQuery(gocb.NewN1qlQuery(req.Query), nil)
		if err != nil {
			c.JSON(http.StatusNotFound,
				http.StatusText(http.StatusNotFound))
			return
		}

		var row map[string]interface{}
		var data []map[string]interface{}

		for rows.Next(&row) {
			data = append(data, row)
			// need to make row empty or it will just keep re-writing on the same memory while /
			// adding the next element to the same memory so while the added memory to data being /
			// a pointer, is pointing to the same memory address making it appear as the same element /
			// while it's actually just pointing to the same memory address
			row = nil //todo: try to find a better way to fix this
		}
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest))
		return
	}

	/**/
}

func errorMessage(err error, context string) error {
	msg := fmt.Sprintf("couch Controller :: %s :: %s", context, err.Error())
	// Send to Sentry
	raven.CaptureErrorAndWait(err, map[string]string{"Controller": msg})
	return errors.New(msg)
}

// todo: using upsert instead of insert
// todo: error handle
// todo: return the data after inserting
