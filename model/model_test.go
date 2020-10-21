package model

import (
	"demo_go/db"
	"demo_go/logger"
	"testing"
	"time"
)

func TestModelProduct(t *testing.T)  {
	// Create
	db.DB.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.DB.First(&product, 1) // 根据整形主键查找
	db.DB.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.DB.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.DB.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.DB.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.DB.Delete(&product, 1)
}


func TestModelUserCreate(t *testing.T)  {
	// Create
	email := "hyhlinux@163.com"
	result := db.DB.Create(&User{Name: "D42", Email: email, Age: 18, Birthday: time.Now()})
	if result.Error != nil {
		logger.Error(result.Error)
	}
	var user User
	db.DB.First(&user, "name = ?", "D42")
	//db.DB.Model(&user).Update("Name", "test")
	logger.Debugf("id:%v name:%v age:%v Birthday:%v", user.ID, user.Name, user.Age, user.Birthday)
	//if user.ID != 0{
	//	logger.Debugf("delete id:%v", user.ID)
	//	db.DB.Delete(&user, user.ID)
	//}
	//result.Error        // 返回 error
	//result.RowsAffected // 返回插入记录的条数
}

func TestModelUserUpdate(t *testing.T)  {
	var user User
	db.DB.First(&user, "name = ?", "D42")
	user.Age = 38
	db.DB.Model(&user).Save(user)
}

func TestModelUserRead(t *testing.T)  {
	var user User
	db.DB.First(&user, "name = ?", "D42")
	logger.Debugf("bf id:%v name:%v age:%v", user.ID, user.Name, user.Age)
}

func TestModelUserReadMany(t *testing.T)  {
	rows, err := db.DB.Model(&User{}).Where("name = ?", "D42").Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	logger.Debugf("rows:%v ", rows)
	for rows.Next() {
		var user User
		// ScanRows 方法用于将一行记录扫描至结构体
		db.DB.ScanRows(rows, &user)
		// 业务逻辑...
		logger.Debugf(" id:%v name:%v age:%v", user.ID, user.Name, user.Age)
	}

}

func TestModelUserReadRaw(t *testing.T)  {
	//users := make([]User, 0)
	//db.DB.Raw("SELECT id, name, age FROM users WHERE name = ?", "D42").Scan(&users)
	//logger.Debugf("users:%v", users)
	var user User
	db.DB.Raw("SELECT id, name, age FROM users WHERE name = ? limit 1", "D42").Scan(&user)
	logger.Debugf("raw sql id:%v name:%v age:%v", user.ID, user.Name, user.Age)
	//for user := range users {
		//logger.Debugf("u:%v", user)
	//	logger.Debugf("raw sql id:%v name:%v age:%v", user.ID, user.Name, user.Age)
	//}
}


func TestModelUserDelete(t *testing.T)  {
	var user User
	db.DB.First(&user, "name = ?", "D42")
	db.DB.Delete(&user, user.ID)
	logger.Debugf("delete id:%v name:%v age:%v", user.ID, user.Name, user.Age)
}
