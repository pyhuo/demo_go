package controller

import (
	"demo_go/db"
	"demo_go/logger"
	"demo_go/model"
	"github.com/gin-gonic/gin"
)

//func pool(w http.ResponseWriter, r *http.Request) {
//	rows, err := db.DB.GET("SELECT * FROM user limit 1")
//	defer rows.Close()
//
//	columns, _ := rows.Columns()
//	scanArgs := make([]interface{}, len(columns))
//	values := make([]interface{}, len(columns))
//	for j := range values {
//		scanArgs[j] = &values[j]
//	}
//
//	record := make(map[string]string)
//	for rows.Next() {
//		//将行数据保存到record字典
//		err = rows.Scan(scanArgs...)
//		for i, col := range values {
//			if col != nil {
//				record[columns[i]] = string(col.([]byte))
//			}
//		}
//	}
//
//	fmt.Println(record)
//	fmt.Fprintln(w, "finish")
//}


func PoolHandler(c *gin.Context) {
	var user model.User
	tx := db.DB.Raw("SELECT id, name, age FROM users WHERE name = ? limit 1", "D42").Scan(&user)
	if tx.Error != nil {
		panic(tx.Error)
	}
	logger.Debugf("raw sql id:%v name:%v age:%v", user.ID, user.Name, user.Age)
	c.JSON(200, user)
}