package bootstrap

import (
	"goblog/pkg/model"
	"time"
)

func SetupDB() {
	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
