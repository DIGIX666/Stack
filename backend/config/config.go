package config

import (
	"fmt"
	"log"
	"os"

	"github.com/DIGIX666/stack/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	if err := models.AutoMigrateModels(db); err != nil {
		log.Fatal("Migration failed:", err)
	}
	DB = db

}

// InitDBWithDialector permet d’injecter un dialector (Postgres, SQLite…)
func InitDBWithDialector(dialector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	// Auto‑migration des modèles
	if err := models.AutoMigrateModels(db); err != nil {
		log.Printf("Migration failed: %v", err)
		return nil, fmt.Errorf("migration failed: %w", err)
	}
	DB = db
	return db, nil
}
