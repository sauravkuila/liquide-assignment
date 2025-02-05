package main

import (
	"liquide-assignment/api"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// start server
	if err := api.Start(); err != nil {
		log.Fatal("Failed to start server, err:", err)
		os.Exit(1)
	}
	addShutdownHook()
}

func addShutdownHook() {
	// when receive interruption from system shutdown server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	log.Println("Quit/Interrupt signal detected. Gracefully closing connections")
	//shutdown server
	api.ShutdownRouter()

	log.Printf("All done! Wrapping up here for PID: %d", os.Getpid())
}
