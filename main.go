package main

import (
	"context"
	"github.com/maribowman/gin-skeleton/app"
	"github.com/maribowman/gin-skeleton/app/config"
	"github.com/maribowman/gin-skeleton/app/model"
	"github.com/maribowman/gin-skeleton/app/repository"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var databaseClient model.DatabaseClient

func init() {
	initLogger()
	var err error
	if databaseClient, err = repository.NewDatabaseClient(); err != nil {
		log.Fatal().Err(err).Msg("unable to init database connection")
	}
}

func initLogger() {
	var logger zerolog.Logger
	if config.Config.Logging.OutputFormat == "TEXT" {
		logFormat := zerolog.ConsoleWriter{Out: os.Stdout}
		logger = log.Output(logFormat).With().Timestamp().Caller().Logger()
	} else {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	}
	level, err := zerolog.ParseLevel(config.Config.Logging.LogLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
	log.Logger = logger
	log.Info().Msgf("logging on %v level", level)
}

func main() {
	server, err := app.InitServer(databaseClient)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init server")
	}
	go func() {
		log.Info().Msgf("running service on port %d", config.Config.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to boot server")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("closing database connection...")
	databaseClient.CloseDatabaseConnections()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("server forced to shutdown")
	}
}
