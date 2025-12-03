package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/garylouisstewart/go-micro/internal/handlers"
	"github.com/garylouisstewart/go-micro/internal/middleware"
)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handlers.HelloWorldHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)

	handlerWithLogging := middleware.LoggingMiddleware(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	svr := &http.Server{
		Addr: ":" + port,
		Handler: handlerWithLogging,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
	}

	go func() {
		middleware.Logger.Info("Starting server on port", "port", port)
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			middleware.Logger.Error("Server failed", "error", err)
	  }
	}()


	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	middleware.Logger.Info("Shutting down server....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		middleware.Logger.Error("Server forced to shutdown,", "error", err)
	}

	middleware.Logger.Info("Server exiting")
}

