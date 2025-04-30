package db

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/project-platypus-perry/v8/internal/config"
	"github.com/project-platypus-perry/v8/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(cfg *config.Config) (*gorm.DB, error) {
	logger.Info("Initializing Postgres")
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Info("Postgres initialized")
	return db, nil
}
