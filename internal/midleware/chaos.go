package midleware

import (
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
			Use for example 1G of RAM

			Simulate broken application using RAM
		*/
		if config.LOG_LEVEL == "debug" {
			log.Println("DEBUG [Chaos]: Doing something wrong 😈...")
		}
		if config.HEAVY_RAM > 0 || c.DefaultQuery("heavy", "") == "ram" {
			if config.LOG_LEVEL == "debug" {
				log.Println("DEBUG [Chaos]: Filling memmory, because you set it by argument ?heavy=ram or option -heavy-ram")
			}
			overflow := make([]byte, 1024*1024*config.HEAVY_RAM)
			for i := 0; i < len(overflow); i += 1024 {
				overflow[i] = byte(i / 42)
			}
			time.Sleep(time.Duration(300))
			for i := 0; i < len(overflow); i += 1024 {
				overflow[i] = byte(i % 102)
			}

		}

		if config.HEAVY_CPU > 0 || c.DefaultQuery("heavy", "") == "cpu" {
			// Simulate CPU heavy task
			if config.LOG_LEVEL == "debug" {
				log.Println("DEBUG [Chaos]: Generating high CPU load because you set it by argument ?heavy=cpu or -heavy-cpu")
			}
			done := make(chan bool)
			go func() {
				// Simulate CPU load for 1 seconds
				end := time.Now().Add(time.Duration(config.HEAVY_CPU) * time.Millisecond)
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
