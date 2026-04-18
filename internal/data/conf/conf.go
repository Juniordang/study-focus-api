package conf

import (
	"fmt"

	"gorm.io/gorm"
)

type Config struct {
	Port   string
	DBPath string
	DNS    string
}

var db *gorm.DB

func Init() error {
	var err error

	db, err = IntializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

func GetSqlite() *gorm.DB {
	return db
}

func Load() Config {
	path := "./db/main.db"

	return Config{
		Port:   ":8080",
		DNS:    path + "?_pragma=foreign_keys(1)",
		DBPath: path,
	}
}
