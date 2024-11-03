package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/pradeepbepari/golang_microservices/cmd"
	"github.com/pradeepbepari/golang_microservices/pkg/config"
	"github.com/pradeepbepari/golang_microservices/pkg/logger"
	"github.com/pradeepbepari/golang_microservices/pkg/tracer"
	"github.com/pradeepbepari/golang_microservices/routes"
	"go.opentelemetry.io/otel"
	"golang.org/x/exp/slog"
)

//go:embed configs
var embedFS embed.FS

func main() {
	//setup the configs
	_config, err := config.LoadConfig(embedFS)
	if err != nil {
		log.Fatal(err)
		return
	}
	//setup Logger
	_loggger := logger.NewLogger()

	//setup tracer
	_tracer, err := tracer.NewTraceProvider()
	if err != nil {
		log.Fatal("failed to start tracing %w", err)
	}
	//shutting down the tracer
	defer func() {
		if err := _tracer.Shutdown(context.Background()); err != nil {
			log.Panicf("Failed to shutdown Tracing %v", err)
		}
	}()
	otel.SetTracerProvider(_tracer)
	database := make(chan *sql.DB)
	server := make(chan *gin.Engine)
	wg := &sync.WaitGroup{}
	root := cmd.NewCommand(&cmd.Server{
		Config:   _config,
		Database: database,
		Server:   server,
		Wg:       wg,
	})
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(database)
		close(server)
	}()
	db := <-database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	}

	slog.Info("Database connection successful")
	router := <-server
	routes.InatiliazeCependencies(db, router, _loggger)
	slog.Info("Http server connection successful")
	if err := router.Run(fmt.Sprintf(":%s", _config.ServerPort)); err != nil {
		log.Fatal(err)
	}
}
