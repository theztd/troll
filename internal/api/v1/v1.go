package v1

import (
	"log"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/handlers"
)

func RoutesAdd(rtG *gin.RouterGroup) {
	r := rtG.Group("/")
	log.Println("Loading V1 routes...")

	r.Use(requestid.New())

	r.GET("/status", handlers.GetStatus)
	r.GET("/info", handlers.GetInfo)
	r.GET("/headers", handlers.GetAllHeaders)
	r.GET("/:item/*id", handlers.GetSlowResponse)
	r.POST("/:item/*id", handlers.GetSlowResponse)
}
