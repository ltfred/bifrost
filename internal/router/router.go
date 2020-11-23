package router

import (
	"bifrost/internal/handler"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.GET("/api/v1/today/", handler.TodayInfo)
	r.GET("/api/v1/history/", handler.History)
	return r
}
