package repository

import (
	"fmt"
	"github.com/go-sanitize/sanitize"
	"github.com/maribowman/gin-skeleton/app/config"
	"github.com/maribowman/gin-skeleton/app/model"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type DatabaseClient struct {
	db *gorm.DB
}

func NewDatabaseClient() (model.DatabaseClient, error) {
	psql, err := initPostgresConnection()
	if err != nil {
		return nil, err
	}
	client := DatabaseClient{
		db: psql,
	}
	return &client, nil
}

func initPostgresConnection() (*gorm.DB, error) {
	timeoutChannel := make(chan *gorm.DB, 1)
	defer close(timeoutChannel)

	go func() {
		log.Info().Msg("establishing postgres connection")
		postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s",
			config.Config.Database.Postgres.Host, config.Config.Database.Postgres.Port, config.Config.Database.Postgres.User, config.Config.Database.Postgres.Password, config.Config.Database.Postgres.Name, config.Config.Database.Postgres.SslMode, config.Config.Database.Postgres.TimeZone)
		psql, err := gorm.Open(postgres.Open(postgresInfo), &gorm.Config{
			//Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatal().Err(err)
		} else {
			timeoutChannel <- psql
		}
	}()

	var psql *gorm.DB
	select {
	case <-time.After(5 * time.Second):
		log.Fatal().Msg("database connection-init timed out - shutting down.")
	case psql = <-timeoutChannel:
		log.Info().Msg("opened postgres connection")
	}

	log.Info().Msg("testing database connection") // -> db connection is lazily loaded
	db, _ := psql.DB()
	if err := db.Ping(); err != nil {
		log.Error().Err(err).Msg("could not ping database")
		return nil, err
	}

	//log.Info().Msg("setting up relations")
	//if err = postgres.Migrator().AutoMigrate(); err != nil {
	//	return nil, err
	//}

	log.Info().Msg("setting up prometheus interface")
	// TODO setup prometheus properly

	log.Info().Msg("wiping sensitive data from memory")
	sanitizer, _ := sanitize.New()
	if err := sanitizer.Sanitize(&config.Config.Database); err != nil {
		log.Warn().Err(err).Msg("failed to sanitize database connection attributes")
	}
	return psql, nil
}

func (client *DatabaseClient) CloseDatabaseConnections() {
	postgresDB, err := client.db.DB()
	if err != nil {
		log.Warn().Err(err).Msg("unable to get database connection")
	}
	if err = postgresDB.Close(); err != nil {
		log.Warn().Err(err).Msg("unable to close database connection")
	}
}
