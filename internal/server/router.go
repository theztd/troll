package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiV1 "gitlab.com/theztd/troll/internal/api/v1"
	apiV2 "gitlab.com/theztd/troll/internal/api/v2"
	"gitlab.com/theztd/troll/internal/config"
	"gitlab.com/theztd/troll/internal/handlers"
	"gitlab.com/theztd/troll/internal/midleware"
)

func InitRoutes() *gin.Engine {
	/*
		Please check
		https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
		for details.

	*/
	// router.TrustedPlatform = "X-CDN-IP"
	// router.SetTrustedProxies([]string{"127.0.0.1"})

	router := gin.New()
	router.Use(midleware.Chaos())
	router.Use(midleware.ServerReceivedHeaders())

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
	router.GET("/_healthz/ready.json", handlers.Ready)

	// get global Monitor object
	config.Metrics.SetMetricPath("/_healthz/metrics")
	config.Metrics.Use(router)

	// Websockets endpoint
	router.GET("/ws", handlers.WebsocketRoute)
	router.GET("/websocket", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ws.html", nil)
	})

	// define default for not found
	router.NoRoute(handlers.HandleNotFound)

	v1 := router.Group("v1")
	config.Metrics.Use(v1)
	apiV1.RoutesAdd(v1)

	v2 := router.Group("v2")
	config.Metrics.Use(v2)
	apiV2.RoutesAdd(v2)

	return router
}
