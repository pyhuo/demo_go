package model

import (
	"demo_go/db"
	"demo_go/logger"
)


func init() {
	// 迁移 schema
	err := db.DB.AutoMigrate(
		&Product{},
		&User{},
		)
	logger.Debugf("ORM AutoMigrate err:%v", err)
	if err != nil {
		panic(err)
	}
}