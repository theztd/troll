/*
Hello world web application (example)
*/
package main

import (
	"flag"
	"fmt"
	"time"

	"gitlab.com/theztd/troll/internal/config"
	"gitlab.com/theztd/troll/internal/libs"
	"gitlab.com/theztd/troll/internal/server"
)

func main() {
	// declare arguments
	flag.StringVar(&config.NAME, "name", "troll", "Define custom application name")
	flag.IntVar(&config.WAIT, "wait", 0, "Minimal wait time before each request")
	flag.StringVar(&config.DOC_ROOT, "root", "./public", "Define document root for serving files")
	flag.StringVar(&config.V2_PATH, "v2-path", "./v2_api.yaml", "Define path to v2 api endpoint configuration yaml")
	flag.IntVar(&config.FAIL_FREQ, "fail", 0, "Returns 503. Set 1 - 10, where 10 = 100% error rate.")
	flag.IntVar(&config.FILL_RAM, "fill-ram", 0, "Fill ram with each request. Set number in bytes.")
	flag.IntVar(&config.READY_DELAY, "ready-delay", 0, "Simulate long application init (seconds).")

	flag.Parse()

	// it is better to be configurable via env
	config.ADDRESS = libs.GetEnv("ADDRESS", ":8080")
	config.LOG_LEVEL = libs.GetEnv("LOG_LEVEL", "info")

	if config.READY_DELAY > 0 {
		fmt.Printf("Application init")
		for i := 0; i < config.READY_DELAY; i++ {
			time.Sleep(time.Duration(1 * time.Second))
			fmt.Printf(".")
		}
		fmt.Printf(" DONE\n\n")

	}

	// It is enought
	router := server.InitRoutes()

	router.Run(config.ADDRESS)
}
