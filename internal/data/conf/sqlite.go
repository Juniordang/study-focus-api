package conf

import (
	"os"

	"github.com/Juniordang/study-focus-api/cmd/config"
	"github.com/Juniordang/study-focus-api/internal/data/schema"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func IntializeSQLite() (*gorm.DB, error) {
	logger := config.GetLogger("sqlite")

	dbConf := Load()

	_, err := os.Stat(dbConf.DBPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbConf.DBPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbConf.DNS),
		&gorm.Config{})
	if err != nil {
		logger.Errf("sqlite opening error %v", err)
		return nil, err
	}

	err = db.AutoMigrate(
		&schema.Flashcard{},
		&schema.Discipline{},
		&schema.StudySession{},
		&schema.User{},
	)
	if err != nil {
		logger.Errf("sqlite automigration error: %v", err)
	}

	return db, nil

}
