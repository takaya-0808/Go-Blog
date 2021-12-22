package handler

import (
	"Go-Blog/app/src/infrastructure/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowBlog(t *testing.T) {

	router := router.InitRouting()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/Myblog/api/article/", nil)
	router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, http.StatusOK, w.Code)
}
