package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tauki/go-practice/couchbase/save-to-couchbase/models"
	"github.com/tauki/go-practice/couchbase/save-to-couchbase/router/middleware"
	"net/http"
)

func InitRouter(cfg *models.Config) (*gin.Engine, error) {
	router := NewRouter()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Route Not Found"})
	})

	InitCouchDBRouter(router, cfg)

	return router, nil
}

func NewRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	middleware.InitiateMiddleware(router)
	return router
}
