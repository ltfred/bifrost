package test

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestHistory(t *testing.T) {
	resp := Get("/api/v1/history/", nil, t)
	assert.Equal(t, resp.Code, 200)
}
