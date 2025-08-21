package v1

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/handlers"
)

func RoutesAdd(rtG *gin.RouterGroup) {
	r := rtG.Group("/")

	r.Use(requestid.New())
	r.GET("/headers", handlers.GetAllHeaders)
	r.GET("/:item/*id", handlers.GetSlowResponse)
	r.POST("/:item/*id", handlers.GetSlowResponse)
}
