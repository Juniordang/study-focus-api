package conf

import (
	"github.com/Juniordang/study-focus-api/cmd/config"
	"github.com/Juniordang/study-focus-api/internal/data/schema"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func IntializeSQLite() (*gorm.DB, error) {
	logger := config.GetLogger("sqlite")

	dbConf := Load()

	db, err := gorm.Open(sqlite.Open(dbConf.DNS),
		&gorm.Config{})
	if err != nil {
		logger.Errf("sqlite opening error %v", err)
		return nil, err
	}

	err = db.AutoMigrate(
		&schema.Flashcard{},
		&schema.Schedule{},
		&schema.StudySession{},
		&schema.User{},
		&schema.AIHistory{},
		&schema.StudyHistory{},
		&schema.Subject{},
	)
	if err != nil {
		logger.Errf("sqlite automigration error: %v", err)
	}

	return db, nil

}
