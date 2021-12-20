package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_Create(t *testing.T) {
	router := gin.Default()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/levels", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestController_Delete(t *testing.T) {
	router := gin.Default()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/levels/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestController_FindAll(t *testing.T) {
	router := gin.Default()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/levels", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestController_FindOne(t *testing.T) {
	router := gin.Default()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/levels/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestController_Update(t *testing.T) {
	router := gin.Default()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/levels/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
