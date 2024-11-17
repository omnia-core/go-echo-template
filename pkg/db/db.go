package db

import (
	"fmt"
	"log"
	"time"

	"github.com/omnia-core/go-echo-template/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresqlDB(cfg *config.Config) *gorm.DB {
	dataSourceFormat := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"

	dsn := fmt.Sprintf(
		dataSourceFormat,
		cfg.Postgresql.Host,
		cfg.Postgresql.Port,
		cfg.Postgresql.User,
		cfg.Postgresql.Password,
		cfg.Postgresql.DBName,
	)

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN: dsn,
		}), &gorm.Config{
			TranslateError: true,
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
		},
	)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(cfg.Postgresql.Options.Connections)
	sqlDB.SetMaxOpenConns(cfg.Postgresql.Options.Connections)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
