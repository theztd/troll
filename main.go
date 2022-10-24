/*
  Hello world web application (example)
*/
package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// declare arguments
	flag.StringVar(&NAME, "name", "troll", "Define custom application name")
	flag.IntVar(&WAIT, "wait", 0, "Minimal wait time before each request")
	flag.StringVar(&DOC_ROOT, "root", "./public", "Define document root for serving files")
	flag.StringVar(&V2_PATH, "v2-path", "./v2_api.yaml", "Define path to v2 api endpoint configuration yaml")
	flag.Parse()

	// it is better to be configurable via env
	ADDRESS = getEnv("ADDRESS", ":8080")

	// It is enought
	getRoutes()
	router.Run(ADDRESS)
}

func getRoutes() {
	/*
		Please check
		https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
		for details.

	*/
	// router.TrustedPlatform = "X-CDN-IP"
	// router.SetTrustedProxies([]string{"127.0.0.1"})

	// register static dir
	router.Static("/public", DOC_ROOT)
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Troll",
			"message": "Application that helps you with mocking, generating slow responses etc.",
		})
	})

	router.GET("/404", dumpRequest)

	v1 := router.Group("v1")
	v1RoutesAdd(v1)

	v2 := router.Group("v2")
	v2RoutesAdd(v2)

}
