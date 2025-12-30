package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/backend"
	"gitlab.com/theztd/troll/internal/config"
)

func PSQLRoute(r *gin.RouterGroup, ep config.Endpoint) {

	r.GET(ep.Path, func(ctx *gin.Context) {
		psql, err := backend.NewPSQL(ep.DSN)
		if err != nil {
			log.Println("PGERR [conn]:", err)
		}

		results, err := psql.RunAndRenderTpl(ep.Query, ep.Response)
		if err != nil {
			log.Println("PGERR [exec]:", err)
		}

		ctx.JSON(http.StatusOK, json.RawMessage(results))

	})

}
