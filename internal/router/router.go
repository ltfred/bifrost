package router

import (
	"bifrost/internal/handler"
	"bifrost/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": time.Now().Format(pkg.DateTimeFormat)})
	})
	r.GET("/api/v1/today/", handler.TodayInfo)
	r.GET("/api/v1/history/", handler.History)
	r.GET("/api/v1/district/", handler.District)
	return r
}
