package midleware

import (
	"fmt"
	"log"
	"net/http/httputil"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/config"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

func indentLines(s, prefix string) string {
	lines := strings.Split(s, "\n")
	for i := range lines {
		lines[i] = prefix + lines[i]
	}
	return strings.Join(lines, "\n")
}

func AuditLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		reqDump, _ := httputil.DumpRequest(c.Request, true)
		var lineColor string
		if c.Writer.Status() < 399 {
			lineColor = Green
		} else {
			lineColor = Red
		}
		if config.LOG_LEVEL == "debug" {
			log.Printf("%sINFO [AuditLog]: %-5s %-50s %-5d From: %s%s", lineColor, c.Request.Method, c.Request.RequestURI, c.Writer.Status(), c.ClientIP(), Reset)
			fmt.Println(indentLines(string(reqDump), "  > "))
		} else {
			log.Printf("%sINFO [AuditLog]: %-5s %-50s %-5d From: %s UA: %.20s%s", lineColor, c.Request.Method, c.Request.RequestURI, c.Writer.Status(), c.ClientIP(), c.Request.UserAgent(), Reset)
		}
	}
}
