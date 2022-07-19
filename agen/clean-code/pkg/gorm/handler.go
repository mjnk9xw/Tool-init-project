package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dbTypePostgresql = "psql"
	dbTypeMysql      = "mysql"
	dbTypeSqlite     = "sqlite"
)

type Config struct {
	DatabaseDSN  string `json:"database_dns" mapstructure:"database_dns"`
	DatabaseType string `json:"database_type" mapstructure:"database_type"`
}

func New(cfg *Config) (*gorm.DB, error) {
	// TODO(latenssi): safeguard against nil config?

	var dialector gorm.Dialector
	switch cfg.DatabaseType {
	default:
		panic(fmt.Sprintf("database type '%s' not supported", cfg.DatabaseType))
	case dbTypePostgresql:
		dialector = postgres.Open(cfg.DatabaseDSN)
	case dbTypeMysql:
		dialector = mysql.Open(cfg.DatabaseDSN)
	case dbTypeSqlite:
		dialector = sqlite.Open(cfg.DatabaseDSN)
	}

	options := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	db, err := gorm.Open(dialector, options)
	if err != nil {
		return &gorm.DB{}, err
	}

	return db, nil
}

func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("unable to close database")
	}

	if err := sqlDB.Close(); err != nil {
		panic("unable to close database")
	}
}
