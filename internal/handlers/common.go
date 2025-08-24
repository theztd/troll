package handlers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/config"
)

func DumpRequest(c *gin.Context) {
	reqDump, _ := httputil.DumpRequest(c.Request, true)
	urlParams := c.Request.URL.RawQuery
	referer := c.Request.Referer()
	userAgent := c.Request.UserAgent()
	reqId := requestid.Get(c)
	headers := strings.Split(string(reqDump), "\r\n")
	data := ""
	message := "Returns received body"

	// Ready raw body (limit 1 MiB)
	body, err := io.ReadAll(io.LimitReader(c.Request.Body, 1<<20))
	if err != nil {
		log.Println("ERR [DumpRequest]: Unable to get request body.")
		message = "Unable to get request body."
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	data = string(body)

	c.JSON(http.StatusOK, gin.H{
		"message":    message,
		"version":    config.VERSION,
		"app_name":   "troll",
		"node":       config.HOSTNAME,
		"client_ip":  c.ClientIP(),
		"referer":    referer,
		"user-agent": userAgent,
		"reqId":      reqId,
		"headers":    headers,
		"url_params": urlParams,
		"data":       data,
	})

	if config.LOG_LEVEL == "debug" {
		log.Println("DEBUG [DumpRequest]: Received url params.")
		fmt.Println(urlParams)
		log.Println("DEBUG [DumpRequest]: Received headers.")
		fmt.Println(headers)
		log.Println("DEBUG [DumpRequest]: Received body.")
		fmt.Println(data)
	}
}

func Ready(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "pong",
		"node": config.HOSTNAME,
	})
}

func GetAllHeaders(c *gin.Context) {
	reqDump, _ := httputil.DumpRequest(c.Request, true)
	if config.LOG_LEVEL == "debug" {
		log.Println("DEBUG [GetAllHeaders]: Dump all headers.")
		fmt.Println(string(reqDump))
	}
	c.JSON(http.StatusOK, gin.H{
		"rec_headers": strings.Split(string(reqDump), "\r\n"),
	})
}

func HandleNotFound(c *gin.Context) {
	reqDump, _ := httputil.DumpRequest(c.Request, true)
	if config.LOG_LEVEL == "debug" {
		log.Println("DEBUG [HandleNotFound]: Dump request for debug")
		fmt.Println(string(reqDump))
	}
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"message": "You are looking for something what we are looking for too... Contact us and lets try to find it together :-)",
	})
}
