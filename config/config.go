package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("failed to initialize SQLite database: %w", err)
	}
	return nil
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}

func GetSQLite() *gorm.DB {
	return db
}
