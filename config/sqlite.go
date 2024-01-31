package config

import (
	"gobooks/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitiSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/books.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errf("sqlite opening error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Book{}, &schemas.FinishedBook{})
	if err != nil {
		logger.Errf("sqlite auto migration error: %v", err)
		return nil, err
	}
	return db, nil
}
