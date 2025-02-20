package infrastructure

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSetupDatabase(t *testing.T) {
	originalEnv := os.Getenv("GO_ENV")
	originalDBHost := os.Getenv("DB_HOST")

	defer func() {
		os.Setenv("GO_ENV", originalEnv)
		os.Setenv("DB_HOST", originalDBHost)
		os.Remove("test.db")
		os.Remove("personalfinance.db")
	}()

	t.Run("should setup test database correctly", func(t *testing.T) {
		os.Setenv("GO_ENV", "test")

		SetupDatabase()

		assert.NotNil(t, DB)
		assert.IsType(t, &gorm.DB{}, DB)

		_, err := os.Stat("test.db")
		assert.NoError(t, err)
	})

	t.Run("should setup development database correctly", func(t *testing.T) {
		os.Setenv("GO_ENV", "development")

		SetupDatabase()

		assert.NotNil(t, DB)
		assert.IsType(t, &gorm.DB{}, DB)

		_, err := os.Stat("personalfinance.db")
		assert.NoError(t, err)
	})

	t.Run("should setup development database when GO_ENV is empty", func(t *testing.T) {
		os.Setenv("GO_ENV", "")

		SetupDatabase()

		assert.NotNil(t, DB)
		assert.IsType(t, &gorm.DB{}, DB)

		_, err := os.Stat("personalfinance.db")
		assert.NoError(t, err)
	})

	t.Run("should panic when production environment variables are not set", func(t *testing.T) {
		os.Setenv("GO_ENV", "production")
		os.Setenv("DB_HOST", "")
		os.Setenv("DB_USER", "")
		os.Setenv("DB_PASSWORD", "")
		os.Setenv("DB_NAME", "")
		os.Setenv("DB_PORT", "")

		assert.Panics(t, func() {
			SetupDatabase()
		})
	})
}

func TestSetupTestDB(t *testing.T) {
	db, err := setupTestDB()

	assert.NoError(t, err)
	assert.NotNil(t, db)

	os.Remove("test.db")
}

func TestSetupDevelopmentDB(t *testing.T) {
	db, err := setupDevelopmentDB()

	assert.NoError(t, err)
	assert.NotNil(t, db)

	os.Remove("personalfinance.db")
}

func TestSetupProductionDB(t *testing.T) {
	originalVars := map[string]string{
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_USER":     os.Getenv("DB_USER"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_NAME":     os.Getenv("DB_NAME"),
		"DB_PORT":     os.Getenv("DB_PORT"),
	}

	defer func() {
		for k, v := range originalVars {
			os.Setenv(k, v)
		}
	}()

	t.Run("should return error when environment variables are not set", func(t *testing.T) {
		os.Setenv("DB_HOST", "")
		os.Setenv("DB_USER", "")
		os.Setenv("DB_PASSWORD", "")
		os.Setenv("DB_NAME", "")
		os.Setenv("DB_PORT", "")

		_, err := setupProductionDB()

		assert.Error(t, err)
	})
}
