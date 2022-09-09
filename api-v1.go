package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func slowResponse(c *gin.Context) {

	// make response randomly slower
	delay := time.Millisecond * time.Duration(WAIT+rand.Intn(500))
	time.Sleep(delay)

	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(data))

	c.JSON(http.StatusOK, gin.H{
		"item":         c.Param("item"),
		"id":           c.Param("id"),
		"reqId":        requestid.Get(c),
		"delay":        delay,
		"receivedData": string(data),
	})
}

func v1RoutesAdd(rtG *gin.RouterGroup) {
	r := rtG.Group("/")
	log.Println("Loading V1 routes...")

	r.Use(requestid.New())

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":   "pong",
			"reqId": requestid.Get(c),
		})
	})

	r.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version":    "0.1.1",
			"app_name":   "troll",
			"client_ip":  c.ClientIP(),
			"referer":    c.Request.Referer(),
			"user-agent": c.Request.UserAgent(),
			"reqId":      requestid.Get(c),
		})
	})

	r.GET("/:item/*id", slowResponse)
	r.POST("/:item/*id", slowResponse)
}
