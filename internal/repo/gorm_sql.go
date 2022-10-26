package repo

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	defaultSqlitePath = "./database.db"
	sqlitePathEnv     = "DB_PATH"
)

// Migrate perform database migrations
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Recipe{})
}

// Connect to the database
// If the database does not exist it will be created
func Connect() (*gorm.DB, error) {
	dbPath := os.Getenv(sqlitePathEnv)
	if dbPath == "" {
		dbPath = defaultSqlitePath
	}

	return gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
}
