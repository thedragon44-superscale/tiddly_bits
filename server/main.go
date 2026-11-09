package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"emv_system/server/transport"
)

func main() {
	log.Println("Initializing EMV system server...")

	// Initialize the central WebSocket hub
	hub := transport.NewWebSocketHub()
	go hub.Run()

	// Register HTTP routes
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		transport.HandleWebSocket(hub, w, r)
	})

	serverAddr := ":8080"
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	// Channel to listen for interrupt signals for graceful shutdown
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("EMV WebSocket server listening on %s\n", serverAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v\n", err)
		}
	}()

	// Block until shutdown signal is received
	<-shutdownChan
	log.Println("Shutting down server gracefully...")
}
