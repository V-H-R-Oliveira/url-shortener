package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url-shortener/1.0/config"
	"url-shortener/1.0/utils"
)

var logger = utils.CreateLogger("[server]> ")

func CreateServer(handler http.Handler, port string) *http.Server {
	return &http.Server{
		Addr:         net.JoinHostPort("", port),
		Handler:      handler,
		ReadTimeout:  config.SERVER_READ_TIMEOUT,
		WriteTimeout: config.SERVER_WRITE_TIMEOUT,
	}
}

func ListenServer(server *http.Server) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Printf("Starting http server at %s\n", server.Addr)

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatal("Failed to listen and serve due error:", err)
		}
	}()

	<-signalChannel
	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Println("Shutting down server")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Failed to gracefully shutdown the server due error:", err)
	}
}
