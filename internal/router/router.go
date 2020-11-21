package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.GET("/api/v1/time", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": time.Now().Format("2006-01-02 15:04:05")})
	})
	return r
}
