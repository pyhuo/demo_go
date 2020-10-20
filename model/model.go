package model

import (
	"demo_go/db"
	"demo_go/logger"
)


func init() {
	// 迁移 schema
	logger.Debug("ORM AutoMigrate")
	db.DB.AutoMigrate(&Product{})
	db.DB.AutoMigrate(&User{})
}