package gormtestutil

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormTest(models ...interface{}) *gorm.DB {
	// Use in-memory SQLite for tests
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to in-memory sqlite database")
	}

	// Enable foreign keys (SQLite disables them by default)
	db.Exec("PRAGMA foreign_keys = ON;")

	// Set a 10-second database lock timeout to avoid "database is locked" errors during tests
	db.Exec("PRAGMA busy_timeout = 10000;")
	// Set journal mode to WAL for better concurrency in tests
	db.Exec("PRAGMA journal_mode = WAL;")

	// Auto-migrate all provided models
	if len(models) > 0 {
		if err := db.AutoMigrate(models...); err != nil {
			panic("failed to migrate test models")
		}
	}

	return db
}
