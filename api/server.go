package api

import (
	"context"
	"liquide-assignment/pkg/service"
	"log"
	"net/http"
	"time"
)

var srv *http.Server
var ctx context.Context

func Start() error {
	ctx = context.Background()

	serviceObj := service.NewServiceGroupObject()
	startRouter(serviceObj)
	return nil
}

func startRouter(obj service.ServiceGroupLayer) {
	srv = &http.Server{
		Addr:    ":8080",
		Handler: getRouter(obj),
	}
	// run api router
	log.Println("starting router")
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server %s", err.Error())
		}
	}()
}

func ShutdownRouter() {
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	log.Println("Shutting down router START")
	defer log.Println("Shutting down router END")
	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Fatalf("Server forced to shutdown. Error: %s", err.Error())
	}
	// catching ctx.Done(). timeout of 2 seconds.
	select {
	case <-timeoutCtx.Done():
		log.Println("timeout of 2 seconds.")
	}
}
