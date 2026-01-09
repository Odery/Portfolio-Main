package main

import (
	"Portfolio"
	"Portfolio/internal"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Metadata about the current build.
// To be injected by the build process.
// TODO: use in logs
var (

	// GitCommitSHA current Git commit SHA of the build, "dev" by default.
	GitCommitSHA = "dev"

	// GitBranch current Git branch of the build, "dev" by default.
	GitBranch = "dev"
)

func main() {
	srv := internal.NewServer("8080", "Portfolio", Portfolio_Main.FS)

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
