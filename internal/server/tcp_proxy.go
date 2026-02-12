package server

import (
	"fmt"
	"log/slog"
	"math/rand"
	"net"
	"time"

	"gitlab.com/theztd/troll/internal/config"
)

func TcpProxyJitter(listenAddr, targetAddr string, minDelay, maxDelay, errRate int) {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		slog.Error("TcpProxy:", err)
	}
	slog.Info(fmt.Sprintf("TcpProxy: Listening on %s and serving content from %s", listenAddr, targetAddr))

	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("TcpProxy: Unable to accept connection", err)
			continue
		}
		go func(clientConn net.Conn) {
			start := time.Now()
			status := "ok"

			defer func() {
				clientConn.Close()
				duration := time.Since(start).Seconds()
				_ = config.Metrics.GetMetric("tcp_proxy_duration_seconds").Observe([]string{status}, duration)
				_ = config.Metrics.GetMetric("tcp_proxy_requests_total").Inc([]string{status})
			}()

			// Establish connection to target
			tgtConn, err := net.Dial("tcp", targetAddr)
			if err != nil {
				slog.Error("TcpProxy: Unable to establish connection to target", err)
				status = "error"
				return
			}
			defer tgtConn.Close()

			// bi directional chan
			done := make(chan struct{}, 2)

			delay := time.Duration(rand.Int63n(int64(maxDelay-minDelay))) * time.Millisecond
			go func() {
				// Client -> Server
				delayCopy(clientConn, tgtConn, delay)
				done <- struct{}{}
			}()

			// Generate random response error by defined TCP_ERROR_RATE
			if rand.Intn(10) < errRate {
				status = "error"
				return
			}

			go func() {
				// Client <- Server
				delayCopy(tgtConn, clientConn, delay)
				done <- struct{}{}
			}()

			<-done

		}(conn)
	}
}

func delayCopy(src, dst net.Conn, delay time.Duration) {
	buffer := make([]byte, 32*1024) // 32KB buffer
	for {
		// Read received data
		n, err := src.Read(buffer)
		if n > 0 {
			time.Sleep(delay)
			// Send received data to target
			_, wErr := dst.Write(buffer[:n])
			if wErr != nil {
				break
			}
		}
		if err != nil {
			break
		}
	}
}
