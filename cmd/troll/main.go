/*
Hello world web application (example)
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"gitlab.com/theztd/troll/internal/config"
	"gitlab.com/theztd/troll/internal/libs"
	"gitlab.com/theztd/troll/internal/server"
)

func main() {
	err := godotenv.Load()
	if err == nil {
		log.Println("INFO: Loading configuration from .env file.")
	}

	// declare arguments
	flag.StringVar(&config.NAME, "name", libs.GetEnv("NAME", "troll"), "Define custom application name")
	flag.IntVar(&config.WAIT, "wait", 0, "Minimal wait time before each request")
	flag.StringVar(&config.DOC_ROOT, "root", libs.GetEnv("DOC_ROOT", "./public"), "Define document root for serving files")
	flag.StringVar(&config.CONFIG_FILE, "config", libs.GetEnv("CONFIG_FILE", "./config.yaml"), "Configure api endpoint")
	flag.StringVar(&config.DSN, "dsn", libs.GetEnv("DSN", ""), "Define database DSN")
	flag.StringVar(&config.ADDRESS, "addr", libs.GetEnv("ADDRESS", ":8080"), "Define address and port where the application listen")
	flag.StringVar(&config.LOG_LEVEL, "log", libs.GetEnv("LOG_LEVEL", "info"), "Define LOG_LEVEL")
	flag.IntVar(&config.FAIL_FREQ, "fail", 0, "Returns 503. Set 1 - 10, where 10 = 100% error rate.")
	flag.IntVar(&config.HEAVY_RAM, "fill-ram", 0, "Fill ram with each request. Set number in bytes.")
	flag.IntVar(&config.HEAVY_CPU, "fill-cpu", 0, "Generate stress on CPU with each request. Set duration in miliseconds (it also works as a delay for request)")
	flag.IntVar(&config.READY_DELAY, "ready-delay", 3, "Simulate long application init (seconds).")

	flag.Parse()

	if config.READY_DELAY > 0 {
		fmt.Printf("Starting application, give me %d sec.", config.READY_DELAY)
		for i := 0; i < config.READY_DELAY; i++ {
			time.Sleep(time.Duration(1 * time.Second))
			fmt.Printf(".")
		}
		fmt.Printf(" DONE\n\n")

	}

	// It is enought
	router := server.InitRoutes()

	fmt.Printf("\n\nAvailable routes:\n")
	for _, r := range router.Routes() {
		fmt.Printf("  ▶︎ %-6s %-30s\n", r.Method, r.Path)
	}
	fmt.Printf("\n\n")
	log.Printf("INFO: Running in mode: \"%s\" and listening on address %s. 😈 Enjoy!", config.LOG_LEVEL, config.ADDRESS)
	router.Run(config.ADDRESS)
}
