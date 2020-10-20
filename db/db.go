package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"

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

	sqlDB, err := DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(2)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}