package main

import (
	"fmt"
	"io/ioutil"
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

func getRoutes() {
	/*
		Build routes tree
	*/

	/*
		Please check
		https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
		for details.

	*/
	// router.TrustedPlatform = "X-CDN-IP"
	// router.SetTrustedProxies([]string{"127.0.0.1"})

	v1 := router.Group("v1")
	v1RoutesAdd(v1)

	v2 := router.Group("v2")
	v2RoutesAdd(v2)

	//plain := router.Group("plain")

}
