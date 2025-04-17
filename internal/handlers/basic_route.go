package handlers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/config"
)

func BasicRoute(r *gin.RouterGroup, ep config.Endpoint) {
	switch ep.Method {
	case "GET":
		r.GET(ep.Path, func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"reqId": requestid.Get(c),
				"msg":   ep.Response,
			})
		})

	case "POST":
		r.POST(ep.Path, func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"reqId": requestid.Get(c),
				"msg":   ep.Response,
			})
		})

	default:
		log.Println("Skip, because method has not been defined " + ep.Path)

	}
}
