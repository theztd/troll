/*
Hello world web application (example)
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/config"
	apiV1 "gitlab.com/theztd/troll/internal/v1"
	apiV2 "gitlab.com/theztd/troll/internal/v2"
)

func jsonLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})

			log["status_code"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["start_time"] = params.TimeStamp.Format("2006/01/02 - 15:04:05")
			log["remote_addr"] = params.ClientIP
			log["response_time"] = params.Latency.String()

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}

func main() {
	// declare arguments
	flag.StringVar(&config.NAME, "name", "troll", "Define custom application name")
	flag.IntVar(&config.WAIT, "wait", 0, "Minimal wait time before each request")
	flag.StringVar(&config.DOC_ROOT, "root", "./public", "Define document root for serving files")
	flag.StringVar(&config.V2_PATH, "v2-path", "./v2_api.yaml", "Define path to v2 api endpoint configuration yaml")
	flag.IntVar(&config.FAIL_FREQ, "fail", 0, "Returns 503. Set 1 - 10, where 10 = 100% error rate.")
	flag.IntVar(&config.FILL_RAM, "fill-ram", 0, "Fill ram with each request. Set number in bytes.")
	flag.IntVar(&config.READY_DELAY, "ready-delay", 0, "Simulate long application init (seconds).")

	flag.Parse()

	// it is better to be configurable via env
	config.ADDRESS = getEnv("ADDRESS", ":8080")
	config.LOG_LEVEL = getEnv("LOG_LEVEL", "info")

	if config.READY_DELAY > 0 {
		fmt.Printf("Application init")
		for i := 0; i < config.READY_DELAY; i++ {
			time.Sleep(time.Duration(1 * time.Second))
			fmt.Printf(".")
		}
		fmt.Printf(" DONE\n\n")

	}

	// It is enought
	router := setRoutes()

	router.Run(config.ADDRESS)
}

func setRoutes() *gin.Engine {
	/*
		Please check
		https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
		for details.

	*/
	// router.TrustedPlatform = "X-CDN-IP"
	// router.SetTrustedProxies([]string{"127.0.0.1"})

	router := gin.New()
	router.Use(MidlewareChaos())

	// register static dir
	router.Static("/public", config.DOC_ROOT)
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Troll",
			"message": "Application that helps you with mocking, generating slow responses etc.",
		})
	})

	// _healthz routes
	router.GET("/_healthz/ready.json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "pass",
			"version": config.VERSION,
			"notes":   "Troll is a very simple webserver returning defined response with configurable delay and a few more features.",
		})
	})

	// router.GET("/_healthz/status.json", HealthDetail)

	// get global Monitor object
	config.Metrics.SetMetricPath("/_healthz/metrics")
	config.Metrics.Use(router)

	// Websockets endpoint
	router.GET("/ws", wsTime)
	router.GET("/websocket", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ws.html", nil)
	})

	// define default for not found
	router.NoRoute(dumpRequest)

	v1 := router.Group("v1")
	config.Metrics.Use(v1)
	apiV1.RoutesAdd(v1)

	v2 := router.Group("v2")
	config.Metrics.Use(v2)
	apiV2.RoutesAdd(v2)

	return router
}

func dumpRequest(c *gin.Context) {
	reqDump, _ := httputil.DumpRequest(c.Request, true)
	fmt.Println(string(reqDump))
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"message": "You are looking for something what we are looking for too... Contact us and lets try to find it together :-)",
	})
}

func MidlewareChaos() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			Use 1G of RAM

			Simulate broken application using RAM
		*/
		if config.FILL_RAM > 0 {
			fmt.Println("INFO: Filling memmory, because you set it by option -fill-ram")
			overflow := make([]byte, 1024*1024*config.FILL_RAM)
			for i := 0; i < len(overflow); i += 1024 {
				overflow[i] = byte(i / 42)
			}
			time.Sleep(time.Duration(300))
			for i := 0; i < len(overflow); i += 1024 {
				overflow[i] = byte(i % 102)
			}

		}

		if c.DefaultQuery("heavy", "") == "cpu" {
			// Simulate CPU heavy task
			log.Println("INFO: Generating high CPU load due to ?heavy=cpu")
			done := make(chan bool)
			go func() {
				// Simulate CPU load for 1 seconds
				end := time.Now().Add(1 * time.Second)
				for time.Now().Before(end) {
					_ = rand.Intn(1000) * rand.Intn(1000) // Perform random calculations
				}
				done <- true
			}()
			<-done

		}

		/*
			Generate 503 errors

			Higher FAIL_FREQ value means more errors
		*/
		if rand.Intn(10) < config.FAIL_FREQ {

			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
				"message": "Troll generates random error, because option -fail has been set. Disable it if you don't wnat to see this error again.",
				"status":  503,
			})
		}

	}
}
