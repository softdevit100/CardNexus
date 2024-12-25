package postgres

import (
	"errors"
	"fmt"
	"tcg-games/app/db"
	"tcg-games/app/helpers"
	"tcg-games/app/models"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ErrConnectingToDB = errors.New("error connecting to db")
var ErrRunningMigrations = errors.New("error running migrations")

// implements Store interface
type Postgres struct {
	db *gorm.DB
}

// GetDB returns a pointer to the database
func New() (db.Store, error) {

	dbName := helpers.Env("DB_NAME")
	dbUser := helpers.Env("DB_USER")
	dbPassword := helpers.Env("DB_PASSWORD")
	dbPort := helpers.Env("DB_PORT")
	dbHost := helpers.Env("DB_HOST_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error().Err(err).Msg("error connecting to db")

		return nil, ErrConnectingToDB
	}

	db.Logger = db.Logger.LogMode(logger.Warn)

	err = db.AutoMigrate(
		&models.Card{},
	)

	if err != nil {
		log.Error().Err(err).Msg("error running migrations")

		return nil, ErrRunningMigrations
	}

	return &Postgres{db}, nil
}
