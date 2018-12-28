package middleware

import (
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"github.com/tauki/go-practice/couchbase/save-to-couchbase/router/middleware/cors"
	"github.com/tauki/go-practice/couchbase/save-to-couchbase/router/middleware/sentry"
)

func InitiateMiddleware(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(sentry.Recovery(raven.DefaultClient, true))
	router.Use(cors.Middleware())
}
