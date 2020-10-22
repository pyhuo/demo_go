package model

import (
	"demo_go/db"
	"demo_go/logger"
	"sort"
	"strings"
	"testing"
	"time"
)

func Now() time.Time{
	return time.Now()
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
	// Create
	//email := "hyhlinux@163.com"
	//result := db.DB.Create(&User{Name: "D42", Email: email, Age: 18, Birthday: time.Now()})
	var user User
	user = User{Name: "ScanRowsUser3", Age: 20, Email: "hyhlinux@165.com", Birthday: time.Now()}
	result := db.DB.Save(&user)
	if result.Error != nil {
		logger.Error(result.Error)
	}
	var newUser User
	db.DB.First(&newUser, "name = ?", user.Name)
	//db.DB.Model(&user).Update("Name", "test")
	//logger.Debugf("id:%v name:%v age:%v Birthday:%v", user.ID, user.Name, user.Age, user.Birthday)
	logger.Debugf("id:%v name:%v age:%v Birthday:%v", newUser.ID, newUser.Name, newUser.Age, newUser.Birthday)
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
	//if !reflect.DeepEqual(results, []Result{{Name: "ScanRowsUser2", Age: 10}, {Name: "ScanRowsUser3", Age: 20}}) {
	//	t.Errorf("Should find expected results")
	//}
	//
	//var ages int
	//if err := db.DB.Table("users").Where("name = ? or name = ?", user2.Name, user3.Name).Select("SUM(age)").Scan(&ages).Error; err != nil || ages != 30 {
	//	t.Fatalf("failed to scan ages, got error %v, ages: %v", err, ages)
	//}
	//
	//var name string
	//if err := db.DB.Table("users").Where("name = ?", user2.Name).Select("name").Scan(&name).Error; err != nil || name != user2.Name {
	//	t.Fatalf("failed to scan ages, got error %v, ages: %v", err, name)
	//}
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
