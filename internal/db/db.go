package db

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/gagan-gaurav/v8/internal/config"
	"github.com/gagan-gaurav/v8/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(cfg *config.Config) *gorm.DB {
	logger.Info("Initializing Postgres")
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
	}

	// Register the Gorm schema driver
	// stmts, err := gormschema.New("postgres").Load(&model.User{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(stmts)

	logger.Info("Postgres initialized")
	return db
}
