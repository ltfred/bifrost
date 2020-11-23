package handler

import (
	"bifrost/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func History(c *gin.Context)  {
	client := &http.Client{}
	now := time.Now()
	month := c.DefaultQuery("month", fmt.Sprintf("%d", int(now.Month())))
	day := c.DefaultQuery("day", fmt.Sprintf("%d", now.Day()))
	url := "https://api.jisuapi.com/todayhistory/query" + "?appkey=" + config.JiSuAppKey +"&month=" + month + "&day=" + day
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "出错啦"})
	}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "出错啦"})
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "出错啦"})
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "出错啦"})
	}
	c.JSON(http.StatusOK, gin.H{"message": result["result"]})
}
