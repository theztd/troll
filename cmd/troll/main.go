/*
Hello world web application (example)
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
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
	// Common args
	flag.StringVar(&config.NAME, "name", libs.GetEnv("NAME", "troll"), "Define custom application name. (NAME)")
	flag.IntVar(&config.READY_DELAY, "ready_delay", libs.GetEnvInt("READY_DELAY", 5), "Simulate long application init [sec]. (READY_DELAY)")
	flag.StringVar(&config.CONFIG_FILE, "config", libs.GetEnv("CONFIG_FILE", ""), "Configure api endpoint. (CONFIG_FILE)")
	flag.StringVar(&config.LOG_LEVEL, "log_level", libs.GetEnv("LOG_LEVEL", "info"), "Define LOG_LEVEL")

	// HTTP args
	flag.IntVar(&config.REQUEST_DELAY, "http.req_delay", libs.GetEnvInt("REQUEST_DELAY", 0), "Minimal delay before response on request [miliseconds]. (REQUEST_DELAY)")
	flag.StringVar(&config.DOC_ROOT, "http.root", libs.GetEnv("DOC_ROOT", "./public"), "Define document root for serving files. (DOC_ROOT)")
	//flag.StringVar(&config.DSN, "dsn", libs.GetEnv("DSN", ""), "Define database DSN")
	flag.StringVar(&config.ADDRESS, "http.addr", libs.GetEnv("HTTP_ADDR", ":8080"), "Define address and port where the application listen. (HTTP_ADDR)")
	flag.IntVar(&config.ERROR_RATE, "http.error_rate", libs.GetEnvInt("HTTP_ERROR_RATE", 0), "Returns 503. Set 1 - 10, where 10 = 100% error rate. (HTTP_ERROR_RATE)")
	flag.IntVar(&config.HEAVY_RAM, "http.fill_ram", libs.GetEnvInt("HEAVY_RAM", 0), "Fill ram with each request [bytes]. (HEAVY_RAM)")
	flag.IntVar(&config.HEAVY_CPU, "http.fill_cpu", libs.GetEnvInt("HEAVY_CPU", 0), "Generate stress on CPU with each request. It also works as a delay for request [milisecodns]. (HEAVY_CPU)")

	// TCP args
	flag.StringVar(&config.TCP_ADDRESS, "tcp.addr", libs.GetEnv("TCP_ADDR", ":9999"), "Define address and port where the tcp proxy listens. (TCP_ADDR)")
	flag.StringVar(&config.TCP_DEST_ADDRESS, "tcp.dest_addr", libs.GetEnv("TCP_DEST_ADDR", "127.0.0.1:8080"), "Define address and port where to send tcp proxy requests. (TCP_DEST_ADDR)")
	flag.IntVar(&config.TCP_MIN_DELAY, "tcp.min_delay", libs.GetEnvInt("TCP_MIN_DELAY", 100), "Simulate long response minimal delay [miliseconds]. (TCP_MIN_DELAY)")
	flag.IntVar(&config.TCP_MAX_DELAY, "tcp.max_delay", libs.GetEnvInt("TCP_MAX_DELAY", 5000), "Simulate long response max delay [miliseconds]. (TCP_MAX_DELAY)")
	flag.IntVar(&config.TCP_ERROR_RATE, "tcp.error_rate", libs.GetEnvInt("TCP_ERROR_RATE", 0), "Simulate random error rate.  Set 1 - 10, where 10 = 100% error rate. (TCP_ERROR_RATE)")

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

	if config.ERROR_RATE > 10 {
		slog.Warn("Maximal FAIL_RATE can be 10. Setting value to 10 (it means 100% error rate)")
		config.ERROR_RATE = 10
	}

	if config.TCP_ERROR_RATE > 10 {
		slog.Warn("Maximal TCP_ERROR_RATE can be 10. Setting value to 10 (it means 100% error rate)")
		config.TCP_ERROR_RATE = 10
	}
	// TCP proxy generating delay and random errors
	go server.TcpProxyJitter(config.TCP_ADDRESS, config.TCP_DEST_ADDRESS, config.TCP_MIN_DELAY, config.TCP_MAX_DELAY, config.TCP_ERROR_RATE)

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
