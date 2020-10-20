## gorm
* [官网](https://gorm.io/zh_CN/docs/connecting_to_the_database.html)


```go
type DeletedAt sql.NullTime
type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
}
```

## CURD
* crate
* First(all, 分页)
* Update
* Delete


```go
func TestModelUser(t *testing.T)  {
	// Create
	email := "hyhlinux@163.com"
	db.DB.Create(&User{Name: "D42", Email: &email, Age: 18})
	var user User
	db.DB.First(&user, "name = ?", "D42")
	db.DB.Model(&user).Update("Name", "test")
	logger.Debugf("id:%v name:%v age:%v", user.ID, user.Name, user.Age)
	if user.ID != 0{
		logger.Debugf("delete id:%v", user.ID)
		//db.DB.Delete(&user, user.ID)
	}
}
```
