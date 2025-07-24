package handlers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	adapter "gitlab.com/theztd/troll/internal/adapter"
	"gitlab.com/theztd/troll/internal/config"
)

func SqlRoute(r *gin.RouterGroup, ep config.Endpoint) {
	switch ep.Method {
	case "GET":
		r.GET(ep.Path, func(c *gin.Context) {
			// zatim nepodporuje argumenty
			resp, err := adapter.RunQuery(ep.Query)
			if err != nil {
				log.Println("ERR [SqlRouter]: Unable to finish database query.", err)
			}

			c.JSON(http.StatusOK, gin.H{
				"reqId": requestid.Get(c),
				"msg":   resp,
			})
		})

	case "POST":
		// zatim nepodporuje argumenty
		resp, err := adapter.RunQuery(ep.Query)
		if err != nil {
			log.Println("ERR [SqlRouter]: Unable to finish database query.", err)
		}

		r.POST(ep.Path, func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"reqId": requestid.Get(c),
				"msg":   resp,
			})
		})

	default:
		log.Println("Skip, because method has not been defined " + ep.Path)

	}
}
