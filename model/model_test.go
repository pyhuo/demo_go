package model

import (
	"demo_go/db"
	"demo_go/logger"
	"gorm.io/gorm"
	"sort"
	"strings"
	"testing"
	"time"
)

func Now() time.Time{
	return time.Now()
}

func handlerDBErr(t *testing.T, result *gorm.DB, key string)  {
	if result.Error != nil {
		t.Error(result.Error)
		return
	}
	logger.Debugf("%s RowsAffected:%d", key, result.RowsAffected)
}


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
	//var user User
	//user = User{Name: "ScanRowsUser3", Age: 20, Email: "hyhlinux@165.com", Birthday: time.Now()}
	//result := db.DB.Debug().Create(&user)
	// 指定字段
	// [1.286ms] [rows:1] INSERT INTO `users` (`name`,`age`,`birthday`) VALUES ('ScanRowsUser3',20,'2020-10-22 14:09:56.913')
	//result := db.DB.Debug().Select("Name", "Age", "Birthday").Create(&user)
	//handlerDBErr(t, result)
	//logger.Debugf("id:%d", user.ID)
	//批量插入
	var users = []User{
		{Name: "ScanRowsUser1", Age: 21, Email: "hyhlinux@165.com", Birthday: time.Now()},
		{Name: "ScanRowsUser2", Age: 22, Email: "hyhlinux@165.com", Birthday: time.Now()},
		{Name: "ScanRowsUser3", Age: 23, Email: "hyhlinux@165.com", Birthday: time.Now()},
	}
	result := db.DB.Debug().Create(&users)
	handlerDBErr(t, result, "bulk create")
	for _, user := range users {
		logger.Debugf("id:%d", user.ID)
	}
	// 根据 Map 创建
	result = db.DB.Model(&User{}).Debug().Create(map[string]interface{}{
		"Name": "ScanRowsUser1", "Age": 21, "Email": "hyhlinux@165.com", "Birthday": time.Now(),
	})
	handlerDBErr(t, result, "map create one")

	// 根据 `[]map[string]interface{}{}` 批量插入
	result = db.DB.Model(&User{}).Debug().Create([]map[string]interface{}{
		{"Name": "ScanRowsUser12", "Age": 21, "Email": "hyhlinux@165.com", "Birthday": time.Now()},
		{"Name": "ScanRowsUser13", "Age": 21, "Email": "hyhlinux@165.com", "Birthday": time.Now()},
		{"Name": "ScanRowsUser14", "Age": 21, "Email": "hyhlinux@165.com", "Birthday": time.Now()},
	})
	handlerDBErr(t, result, "map bulk create")
}

func TestScanRows(t *testing.T) {
	user1 := User{Name: "ScanRowsUser1", Age: 1, Birthday: Now()}
	user2 := User{Name: "ScanRowsUser2", Age: 10, Birthday: Now()}
	user3 := User{Name: "ScanRowsUser3", Age: 20, Birthday: Now()}
	tx := db.DB.Save(&user1).Save(&user2).Save(&user3)
	logger.Debugf("%v", tx)

	rows, err := db.DB.
		Table("users").
		Where("name = ? or name = ?", user2.Name, user3.Name).
		Select("name, age, birthday").
		Rows()
	if err != nil {
		t.Errorf("Not error should happen, got %v", err)
	}

	type Result struct {
		Name string
		Age  int
	}
	var results []Result
	for rows.Next() {
		var result Result
		if err := db.DB.ScanRows(rows, &result); err != nil {
			t.Errorf("should get no error, but got %v", err)
		}
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return strings.Compare(results[i].Name, results[j].Name) <= -1
	})
	for _, u := range results {
		logger.Debugf("row:%v", u)
	}
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
	rows, err := db.DB.Model(&User{}).Debug().Where("name = ?", "D42").Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	logger.Debugf("rows:%v ", rows)
	for rows.Next() {
		var user User
		// ScanRows 方法用于将一行记录扫描至结构体
		err = db.DB.ScanRows(rows, &user)
		if err != nil {
			panic(err)
		}
		// 业务逻辑...
		logger.Debugf(" id:%v name:%v age:%v", user.ID, user.Name, user.Age)
	}
}

func TestModelUserReadRaw(t *testing.T)  {
	var user User
	db.DB.Raw("SELECT id, name, age FROM users WHERE name = ? limit 1", "D42").Scan(&user)
	logger.Debugf("raw sql id:%v name:%v age:%v", user.ID, user.Name, user.Age)
}


func TestModelUserDelete(t *testing.T)  {
	var user User
	db.DB.First(&user, "name = ?", "D42")
	db.DB.Debug().Delete(&user, user.ID)
	logger.Debugf("delete id:%v name:%v age:%v", user.ID, user.Name, user.Age)
}
