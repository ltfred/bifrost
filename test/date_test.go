package test

import (
	router2 "bifrost/internal/router"
	"encoding/json"
	"github.com/go-playground/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router = router2.Routers()

func Get(url string, body io.Reader, t *testing.T) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", url, body)
	assert.Equal(t, err, nil)
	router.ServeHTTP(w, req)
	return w
}

func TestTodayInfo(t *testing.T) {
	resp := Get("/api/v1/today/?date=2020-11-23", nil, t)
	assert.Equal(t, resp.Code, 200)
	var response map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, err, nil)
	assert.Equal(t, response["date"], "2020-11-23")
	assert.Equal(t, response["lunarCalender"], "庚子鼠年十月初九日")
	assert.Equal(t, response["weekday"], float64(1))

	resp = Get("/api/v1/today/?date=2020-1123", nil, t)
	assert.Equal(t, resp.Code, 400)
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, err, nil)
	assert.Equal(t, response["message"], "错误的时间格式")
}
