/*
Hello world web application (example)
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
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
	flag.StringVar(&config.NAME, "name", libs.GetEnv("NAME", "troll"), "Define custom application name. (NAME)")
	flag.IntVar(&config.REQUEST_DELAY, "req-delay", libs.GetEnvInt("REQUEST_DELAY", 0), "Minimal delay before response on request [miliseconds]. (REQUEST_DELAY)")
	flag.StringVar(&config.DOC_ROOT, "root", libs.GetEnv("DOC_ROOT", "./public"), "Define document root for serving files. (DOC_ROOT)")
	flag.StringVar(&config.CONFIG_FILE, "config", libs.GetEnv("CONFIG_FILE", "./config.yaml"), "Configure api endpoint. (CONFIG_FILE)")
	flag.StringVar(&config.DSN, "dsn", libs.GetEnv("DSN", ""), "Define database DSN")
	flag.StringVar(&config.ADDRESS, "addr", libs.GetEnv("ADDRESS", ":8080"), "Define address and port where the application listen. (ADDRESS)")
	flag.StringVar(&config.LOG_LEVEL, "log", libs.GetEnv("LOG_LEVEL", "info"), "Define LOG_LEVEL")
	flag.IntVar(&config.FAIL_FREQ, "fail", libs.GetEnvInt("FAIL_FREQ", 0), "Returns 503. Set 1 - 10, where 10 = 100% error rate. (FAIL_FREQ)")
	flag.IntVar(&config.HEAVY_RAM, "fill-ram", libs.GetEnvInt("HEAVY_RAM", 0), "Fill ram with each request [bytes]. (HEAVY_RAM)")
	flag.IntVar(&config.HEAVY_CPU, "fill-cpu", libs.GetEnvInt("HEAVY_CPU", 0), "Generate stress on CPU with each request. It also works as a delay for request [milisecodns]. (HEAVY_CPU)")
	flag.IntVar(&config.READY_DELAY, "ready-delay", libs.GetEnvInt("READY_DELAY", 5), "Simulate long application init [sec]. (READY_DELAY)")

	flag.Parse()

	// Print received os signals
	chanSig := make(chan os.Signal, 1)
	signal.Notify(chanSig,
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
		syscall.SIGHUP, syscall.SIGUSR1, syscall.SIGUSR2,
	)
	go func() {
		for s := range chanSig {
			// Badass mode is disabled
			if !config.BADASS {
				switch s {
				case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
					log.Printf("INFO: Received signal \"%s\".", s)
					signal.Reset(s)
					_ = syscall.Kill(os.Getpid(), s.(syscall.Signal))
				}
			} else {
				log.Printf("INFO: Received signal \"%s\", but I do not do anything ⛄️.", s)
			}
		}
	}()

	config.HOSTNAME, err = os.Hostname()
	if err != nil {
		log.Println("WARN: Unable to get Hostname", err)
		config.HOSTNAME = fmt.Sprintf("%s-%s", config.NAME, uuid.NewString()[:6])
	}

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
	if err := router.Run(config.ADDRESS); err != nil {
		log.Fatal("FATAL: router.Run error...", err)
	}
}
