package api

import (
	"context"
	"liquide-assignment/pkg/db"
	e "liquide-assignment/pkg/errors"
	"liquide-assignment/pkg/service"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

var srv *http.Server
var ctx context.Context
var databases []*gorm.DB

func Start() error {
	ctx = context.Background()

	//error initialization
	e.ErrorInit()

	databases = make([]*gorm.DB, 0)
	postgresConn, err := db.PsqlConnect()
	if err != nil {
		log.Printf("Failed to connect psql database. Error:%s", err.Error())
		return err
	}

	//extend the list of databases if multiple db connections needed
	databases = append(databases, postgresConn)
	dbObj := db.NewDBObject(postgresConn)

	//pass db object to service layer
	serviceObj := service.NewServiceGroupObject(dbObj)
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
