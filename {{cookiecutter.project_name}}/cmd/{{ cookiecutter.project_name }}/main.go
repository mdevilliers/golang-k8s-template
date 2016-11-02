package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	log.Printf("{{ cookiecutter.project_name }} running...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutdown signal received, exiting...")
}
