package model

import (
	"demo_go/db"
	"demo_go/logger"
	"testing"
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


func TestModelUser(t *testing.T)  {
	// Create
	//email := "hyhlinux@163.com"
	//db.Create(&User{Name: "D42", Email: &email, Age: 18})
	var user User
	db.DB.First(&user, 1)
	logger.Debug("id:", user.ID, "name:", user.Name, "age:", user.Age)
}