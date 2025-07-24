package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/config"
	"golang.org/x/exp/rand"
)

func GetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version":    config.VERSION,
		"app_name":   "troll",
		"client_ip":  c.ClientIP(),
		"referer":    c.Request.Referer(),
		"user-agent": c.Request.UserAgent(),
		"reqId":      requestid.Get(c),
	})
}

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":   "pong",
		"reqId": requestid.Get(c),
	})
}

func Ready(c *gin.Context) {
	log.Println("DEBUG [handler.Ready]: Processing.")
	c.JSON(http.StatusOK, gin.H{
		"status":  "pass",
		"version": config.VERSION,
		"notes":   "Troll is a very simple webserver returning defined response with configurable delay and a few more features.",
	})
	if config.LOG_LEVEL == "debug" {
		log.Println("DEBUG [handler.Ready]: Processed.")
	}
}

func GetAllHeaders(c *gin.Context) {
	reqDump, _ := httputil.DumpRequest(c.Request, true)
	fmt.Println(string(reqDump))
	c.JSON(http.StatusOK, gin.H{
		"rec_headers": strings.Split(string(reqDump), "\r\n"),
	})
}

func GetSlowResponse(c *gin.Context) {
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

func HandleNotFound(c *gin.Context) {
	reqDump, _ := httputil.DumpRequest(c.Request, true)
	fmt.Println(string(reqDump))
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"message": "You are looking for something what we are looking for too... Contact us and lets try to find it together :-)",
	})
}
