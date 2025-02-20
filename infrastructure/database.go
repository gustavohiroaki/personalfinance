package infrastructure

import (
	"fmt"
	"os"

	"github.com/gustavohiroaki/personalfinance/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func SetupDatabase() *gorm.DB {
	env := os.Getenv("GO_ENV")

	switch env {
	case "test":
		DB, err = setupTestDB()
	case "production":
		DB, err = setupProductionDB()
	default:
		DB, err = setupDevelopmentDB()
	}

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	err = DB.AutoMigrate(
		&models.Transaction{},
		&models.Fees{},
		&models.CorporateEvent{},
	)

	if err != nil {
		panic(fmt.Sprintf("Failed to auto-migrate: %v", err))
	}

	return DB
}

func setupTestDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
}

func setupDevelopmentDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("personalfinance.db"), &gorm.Config{})
}

func setupProductionDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
