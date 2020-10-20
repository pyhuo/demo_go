package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"gorm.io/driver/sqlite"
	"demo_go/logger"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/demo_go?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	logger.Debugf("db init:%v", DB)
}