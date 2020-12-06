package handler

import (
	"bifrost/gen/entmodels"
	"bifrost/gen/entmodels/district"
	"bifrost/internal/app"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)


func District(c *gin.Context)  {
	ctx := context.Background()
	name := c.DefaultQuery("name", "")
	var dis *entmodels.District
	var err error
	if name != "" {
		dis, err = app.EntClient.District.Query().Where(district.DistrictName(name)).First(ctx)
	} else {
		dis, err = app.EntClient.District.Query().Where(district.ID(1)).First(ctx)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "没有查到该地区信息"})
	}
	c.JSON(http.StatusOK, gin.H{"name": dis.DistrictName, "shorter_name": dis.ShorterName, "car_code": dis.CarCode})
}
