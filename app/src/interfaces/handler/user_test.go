package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"Go-Blog/app/src/infrastructure/router"

	"github.com/stretchr/testify/assert"
)

func TestShow(t *testing.T) {

	router := router.InitRouting()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/MyBlog/api/v1/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, http.StatusBadRequest, w.Code)
	// assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestIndex(t *testing.T) {
	router := router.InitRouting()
	w := httptest.NewRecorder()

}
