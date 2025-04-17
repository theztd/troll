package v2

import (
	"log"
	"os"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/config"
	"gitlab.com/theztd/troll/internal/handlers"
)

func RoutesAdd(rGroup *gin.RouterGroup) {
	r := rGroup.Group("/")
	log.Println("Loading V2 routes...")

	r.Use(requestid.New())

	r.GET("/info", handlers.GetInfo)
	r.GET("/status", handlers.GetStatus)

	// if v2 yaml configuration exists, generate endpoints
	if _, err := os.Stat(config.V2_PATH); err == nil {
		cfg := config.LoadYaml(config.V2_PATH)

		for _, endpoint := range cfg.Endpoints {
			switch endpoint.Kind {
			case "basic":
				handlers.BasicRoute(rGroup, endpoint)

			case "sql":
				handlers.SqlRoute(rGroup, endpoint)

			default:
				log.Printf("Skip, because kind has not been defined %s (%s)", endpoint.Path, endpoint.Kind)

			}

		}
	} else {
		log.Println("ERR: Unable to find file " + config.V2_PATH)
	}
}
