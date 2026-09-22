package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
)

type DB struct {
	*gorm.DB
}

func Connect(url string) (*DB, error) {
	gormDB, err := gorm.Open(
		postgres.Open(url),
		&gorm.Config{
			Logger: NewCtxLogger(),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("gorm connect failed: %w", err)
	}

	if err := gormDB.Use(otelgorm.NewPlugin()); err != nil {
		return nil, fmt.Errorf("failed to init gorm tracing: %w", err)
	}

	return &DB{gormDB}, nil
}

func (d *DB) Ping() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}
	return sqlDB.Ping()
}

func (d *DB) GetDB() *gorm.DB {
	return d.DB
}
