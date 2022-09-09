package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

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
