package main

import (
	"Portfolio/internal"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv := internal.NewServer("8080", "Portfolio")

	//Handle sig term
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-sig
		log.Println("Shutting down the web server...")
		if err := srv.Stop(); err != nil {
			log.Printf("Shutdown error: %v\n", err)
		}
	}()

	log.Printf("Starting the web server on port %s\n", srv.Port)

	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
