package handler

import (
	"bifrost/pkg"
	"github.com/gin-gonic/gin"
	"github.com/nosixtools/solarlunar"
	"net/http"
	"time"
)

func TodayInfo(c *gin.Context) {
	now := time.Now()
	dateStr := c.DefaultQuery("date", now.Format(pkg.DateFormat))
	loc, _ := time.LoadLocation("Local")
	date, err := time.ParseInLocation(pkg.DateFormat, dateStr, loc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "错误的时间格式"})
	}
	lunarCalender := solarlunar.SolarToChineseLuanr(date.Format(pkg.DateFormat))
	c.JSON(
		http.StatusOK,
		gin.H{
			"date": date.Format(pkg.DateFormat),
			"weekday": date.Weekday(),
			"lunarCalender": lunarCalender,
			"yearDay": date.YearDay(),
		},
		)
}
