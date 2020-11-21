package handler

import (
	"bifrost/pkg"
	"github.com/gin-gonic/gin"
	"github.com/nosixtools/solarlunar"
	"net/http"
	"time"
)

func TodayInfo(c *gin.Context)  {
	now := time.Now()
	lunarCalender := solarlunar.SolarToChineseLuanr(now.Format(pkg.DateFormat))
	c.JSON(http.StatusOK, gin.H{"date": now.Format(pkg.DateTimeFormat), "weekday": now.Weekday(), "lunarCalender": lunarCalender})
}
