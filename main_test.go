package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestV1StatusEndpoint(t *testing.T) {
	mockResponse := `{"msg":"pong","reqId":""}`
	r := SetUpRouter()
	r.GET("/v1/status", v1Status)
	req, _ := http.NewRequest("GET", "/v1/status", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

/*
func TestV1SlowEndpoint(t *testing.T) {
	//mockResponse := `{"msg":"pong","reqId":""}`
	r := SetUpRouter()
	r.GET("/v1/:item/*id", slowResponse)

	req, _ := http.NewRequest("GET", "/v1/user/126", nil)
	fmt.Println(req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println("ERROR!!!!!!!!!!")

	//responseData, _ := ioutil.ReadAll(w.Body)
	//fmt.Println(string(responseData))
	// assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
*/
