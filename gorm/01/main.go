package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo ---->数据表
type UserInfo struct {
	Id     uint
	Name   string
	Gender string
	Hobby  string
}

// gorm操作数据库
func main() {
	// 连接mysql数据库
	db, err := gorm.Open("mysql", "root:LIJIAlijia@tcp(127.0.0.1:6446)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	// 创建表，自动迁移（把结构体和数据表进行一一对应）
	db.AutoMigrate(&UserInfo{})

	// 创建记录数据行
	u1 := UserInfo{520, "它是你", "男", "Dance"}
	db.Create(&u1)
	// 查询数据

	// 更新数据

	// 删除数据

}
