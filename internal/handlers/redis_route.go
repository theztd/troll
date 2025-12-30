package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/backend"
	"gitlab.com/theztd/troll/internal/config"
)

func RedisRoute(r *gin.RouterGroup, ep config.Endpoint) {

	r.GET(ep.Path, func(ctx *gin.Context) {
		redis, err := backend.NewRedis(ep.DSN)
		if err != nil {
			log.Println("RedisERR [conn]:", err)
		}

		results, err := redis.RunAndRenderTpl(ep.Query, ep.Response)
		if err != nil {
			log.Println("RedisERR [exec]:", err)
		}

		ctx.JSON(http.StatusOK, json.RawMessage(results))

	})

}
