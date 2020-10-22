## CURD
* [create](https://gorm.io/zh_CN/docs/create.html)


### 1.create


```go
func TestModelUserCreate(t *testing.T)  {
	var user User
	user = User{Name: "ScanRowsUser3", Age: 20, Email: "hyhlinux@165.com", Birthday: time.Now()}
	//result := db.DB.Debug().Create(&user)
	// 指定字段
	// [1.286ms] [rows:1] INSERT INTO `users` (`name`,`age`,`birthday`) VALUES ('ScanRowsUser3',20,'2020-10-22 14:09:56.913')
	result := db.DB.Debug().Select("Name", "Age", "Birthday").Create(&user)
	if result.Error != nil {
		t.Error(result.Error)
		return
	}
	logger.Debugf("id:%d", user.ID)
	logger.Debugf("RowsAffected:%d", result.RowsAffected)
}
```
