package config

import (
	"github.com/arthurlopesr/go-opportunities/constant"
	"github.com/arthurlopesr/go-opportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	logger.Info("Initializing SQLite database connection...")

	if err := ensureDatabaseFile(); err != nil {
		logger.ErrorFormated(constant.ErrCreateDatabase, err)
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(constant.DatabasePath), &gorm.Config{})
	if err != nil {
		logger.ErrorFormated(constant.ErrToConnectDatabase, err)
		return nil, err
	}

	if err := db.AutoMigrate(&schemas.Opening{}); err != nil {
		logger.ErrorFormated(constant.ErrAutoMigrateDatabase, err)
		return nil, err
	}
	logger.Info("SQLite database connection established successfully.")
	return db, nil
}

func ensureDatabaseFile() error {
	if _, err := os.Stat(constant.DatabasePath); os.IsNotExist(err) {
		if err := os.MkdirAll("./db", os.ModePerm); err != nil {
			return err
		}
		file, err := os.Create(constant.DatabasePath)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}
