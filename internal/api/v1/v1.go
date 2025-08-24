package v1

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/handlers"
	"gitlab.com/theztd/troll/internal/midleware"
)

func RoutesAdd(rtG *gin.RouterGroup) {
	r := rtG.Group("/")

	r.Use(requestid.New())
	r.Use(midleware.Chaos())
	r.Use(midleware.RandomDelay)
	r.GET("/headers", handlers.GetAllHeaders)
	r.GET("/:item/*id", handlers.DumpRequest)
	r.POST("/:item/*id", handlers.DumpRequest)
}
