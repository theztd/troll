package midleware

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/config"
)

func GetAllHeaders(c *gin.Context) {
	reqDump, _ := httputil.DumpRequest(c.Request, true)
	if config.LOG_LEVEL == "debug" {
		fmt.Println(string(reqDump))
	}
	c.JSON(http.StatusOK, gin.H{
		"rec_headers": strings.Split(string(reqDump), "\r\n"),
	})
}

// To implement
func ServerReceivedHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Do next before run myself
		c.Next()
		if config.LOG_LEVEL == "debug" {
			log.Printf("DEBUG [middleware]: Received headers %s", c.Request.Header)
		}

		// TODO: Tady bych chtel, aby se prijate hlavicky mohli vracet v odpovedi jako receivedHeaders: ....
		// if true {
		// 	var lastResponse map[string]interface{}
		// 	if err := c.ShouldBindJSON(&lastResponse); err != nil {
		// 		log.Println("DEBUG [middleware]: Nothing to do...", err)
		// 		c.Next()
		// 	} else {
		// 		lastResponse["receivedHeaders"] = c.Request.Header
		// 		c.JSON(c.Writer.Status(), lastResponse)
		// 	}
		// } else {
		// 	log.Println(c.ContentType())
		// }

	}
}
