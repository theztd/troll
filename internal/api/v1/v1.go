package v1

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/config"
)

func slowResponse(c *gin.Context) {
	wait, _ := strconv.Atoi(c.Query("wait"))

	// make response randomly slower
	delay := (time.Duration(config.WAIT+rand.Intn(500)) * time.Millisecond) + time.Duration(wait)*time.Millisecond
	time.Sleep(delay)
	fmt.Println(delay)

	data, _ := io.ReadAll(c.Request.Body)
	fmt.Println(string(data))

	c.JSON(http.StatusOK, gin.H{
		"item":  c.Param("item"),
		"id":    c.Param("id"),
		"reqId": requestid.Get(c),
		"delay": delay,
		//"receivedData": string(data),
	})
}

func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":   "pong",
		"reqId": requestid.Get(c),
	})
}

func getInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version":    "0.1.1",
		"app_name":   "troll",
		"client_ip":  c.ClientIP(),
		"referer":    c.Request.Referer(),
		"user-agent": c.Request.UserAgent(),
		"reqId":      requestid.Get(c),
	})
}

func allHeaders(c *gin.Context) {
	reqDump, _ := httputil.DumpRequest(c.Request, true)
	fmt.Println(string(reqDump))
	c.JSON(http.StatusOK, gin.H{
		"rec_headers": strings.Split(string(reqDump), "\r\n"),
	})
}

func RoutesAdd(rtG *gin.RouterGroup) {
	r := rtG.Group("/")
	log.Println("Loading V1 routes...")

	r.Use(requestid.New())

	r.GET("/status", getStatus)
	r.GET("/info", getInfo)
	r.GET("/headers", allHeaders)
	r.GET("/:item/*id", slowResponse)
	r.POST("/:item/*id", slowResponse)
}
