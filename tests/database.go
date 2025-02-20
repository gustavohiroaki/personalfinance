package tests

import (
	"database/sql"
	"os"

	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"gorm.io/gorm"
)

func PrepareTests() {
	os.Setenv("ENV", "test")
}

func PrepareDB() (*gorm.DB, *sql.DB) {
	dbInstance := infrastructure.SetupDatabase()
	sqlInstance, _ := dbInstance.DB()
	return dbInstance, sqlInstance
}

func OnClose(sqlInstance *sql.DB) {
	if sqlInstance != nil {
		sqlInstance.Close()
	}
	os.Remove("test.db")
}
