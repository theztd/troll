package midleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"math/rand/v2"

	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/config"
)

func Chaos() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			Use 1G of RAM

			Simulate broken application using RAM
		*/
		if config.FILL_RAM > 0 {
			fmt.Println("INFO: Filling memmory, because you set it by option -fill-ram")
			overflow := make([]byte, 1024*1024*config.FILL_RAM)
			for i := 0; i < len(overflow); i += 1024 {
				overflow[i] = byte(i / 42)
			}
			time.Sleep(time.Duration(300))
			for i := 0; i < len(overflow); i += 1024 {
				overflow[i] = byte(i % 102)
			}

		}

		if c.DefaultQuery("heavy", "") == "cpu" {
			// Simulate CPU heavy task
			log.Println("INFO: Generating high CPU load due to ?heavy=cpu")
			done := make(chan bool)
			go func() {
				// Simulate CPU load for 1 seconds
				end := time.Now().Add(1 * time.Second)
				for time.Now().Before(end) {
					_ = rand.IntN(1000) * rand.IntN(1000) // Perform random calculations
				}
				done <- true
			}()
			<-done

		}

		/*
			Generate 503 errors

			Higher FAIL_FREQ value means more errors
		*/
		if rand.IntN(10) < config.FAIL_FREQ {

			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
				"message": "Troll generates random error, because option -fail has been set. Disable it if you don't wnat to see this error again.",
				"status":  503,
			})
		}

	}
}
