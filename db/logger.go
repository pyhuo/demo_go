package db

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var newLogger logger.Interface

func init() {
	newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			//SlowThreshold: time.Second,   // 慢 SQL 阈值
			SlowThreshold: time.Nanosecond,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
}