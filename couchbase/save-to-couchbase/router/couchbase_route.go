package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tauki/go-practice/couchbase/save-to-couchbase/controller"
	"github.com/tauki/go-practice/couchbase/save-to-couchbase/models"
)

func InitCouchDBRouter(router *gin.Engine, cfg *models.Config) {
	dbCtrl, err := controller.GetCouchDBController(cfg)
	if err != nil {
		panic(err) //todo: no panic
	}

	db := router.Group("couch")
	db.GET("", dbCtrl.GetAll)
	db.POST("", dbCtrl.Create)
	db.GET("search", dbCtrl.Search)
	db.GET("id/:id", dbCtrl.GetByID)

	db.POST("query", dbCtrl.Query)
}
