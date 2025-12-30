package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	apiV1 "gitlab.com/theztd/troll/internal/api/v1"
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
	if config.LOG_LEVEL == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	gin.SetMode("release")

	router := gin.New()
	router.Use(midleware.Chaos())
	router.Use(midleware.AuditLog())
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
	router.GET("/_healthz/info", midleware.RandomDelay, handlers.DumpRequest)
	router.GET("/_healthz/alive", midleware.RandomDelay, handlers.DumpRequest)
	router.GET("/_healthz/ready", midleware.RandomDelay, handlers.Ready)

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

	if _, err := os.Stat(config.CONFIG_FILE); err == nil && config.CONFIG_FILE != "" {
		// if CONFIG_FILE exists and not empty, load routes
		cfg := config.LoadYaml(config.CONFIG_FILE)

		// Game route initialization
		handlers.BackendUrls = cfg.Game.Backends
		if tmplBytes, err := os.ReadFile(cfg.Game.TemplatePath); err == nil {
			handlers.GameTemplate = string(tmplBytes)
		}

		if cfg.Game.Route != "" {
			router.GET(cfg.Game.Route, handlers.GameUI)
			log.Println("INFO: Initialize GAME route 🎲 " + cfg.Game.Route)
		} else {
			router.GET("/the-game", handlers.GameUI)
			log.Println("INFO: Initialize GAME route 🎲 " + "/the-game")
		}

		v1 := router.Group("v1")
		log.Println("INFO: Initialize V1 routes 🏗️ ...")
		for _, endpoint := range cfg.Endpoints {
			switch endpoint.Kind {
			case "basic":
				handlers.BasicRoute(v1, endpoint)

			case "psql":
				// handlers.SqlRoute(v1, endpoint)
				handlers.PSQLRoute(v1, endpoint)

			case "redis":
				//
				handlers.RedisRoute(v1, endpoint)

			case "shell":
				// handlers.ShellRoute(v1, endpoint)
				handlers.PSQLRoute(v1, endpoint)

			default:
				log.Printf("WARN: Skip, because kind has not been defined %s (%s)", endpoint.Path, endpoint.Kind)

			}

		}

	} else if config.CONFIG_FILE != "" && os.IsNotExist(err) {
		// if CONFIG_FILE is not empty and does not exists, than FAIL
		log.Printf("FATAL: Given config file \"%s\" does not exists!!!", config.CONFIG_FILE)
		os.Exit(66) // File not found

	} else {
		log.Println("WARN: Config file is not defined, but continue with defaults..")
		log.Println("INFO: Initialize default routes 🏗️  ...")

		v1 := router.Group("v1")
		config.Metrics.Use(v1)
		apiV1.RoutesAdd(v1)
	}

	if config.LOG_LEVEL == "DEBUG" {
		log.Println("DEBUG [router.InitRoutes]: All routes has been initialized")
	}

	return router
}
