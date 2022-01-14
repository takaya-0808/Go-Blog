package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"Go-Blog/app/src/domain/model"
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

	req, _ := http.NewRequest("GET", "/MyBlog/api/v1/name", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NotEqual(t, http.StatusOK, w.Code)
}

// func TestRegister(t *testing.T) {

// 	router := router.InitRouting()
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)
// 	body := strings.NewReader("{\"UserName\":\"senss\", \"UserEmail\":\"e17575s\", \"UserPassWord\":\"passss\"}")
// 	body := strings.NewReader(`{"UserName":"sen", "UserEmail":"e175751", "UserPassWord":"pass"}`)
// 	p := model.RegisterModel{
// 		UserName:     "fugafuga",
// 		UserEmail:    "e175751@ie.u-ryukyu.ac.jp",
// 		UserPassWord: "password",
// 	}
// 	body, _ := json.Marshal(p)
// 	c.Request, _ = http.NewRequest(http.MethodPost, "/MyBlog/api/auth/register", bytes.NewBuffer(body))
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Accept", "application/json")

// 	router.ServeHTTP(w, c.Request)
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.NotEqual(t, http.StatusBadRequest, w.Code)
// }

func TestLogin(t *testing.T) {

	router := router.InitRouting()
	w := httptest.NewRecorder()

	form := &model.LoginModel{
		UserName:     "sen",
		UserPassWord: "pass",
	}
	body, _ := json.Marshal(form)

	// bodys := bytes.NewBufferString("{\"UserName\":\"sen\",  \"UserPassWord\":\"pass\"}")
	req, _ := http.NewRequest(http.MethodPost, "/MyBlog/api/auth/login", bytes.NewBuffer(body))

	// req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, http.StatusInternalServerError, w.Code)
}
