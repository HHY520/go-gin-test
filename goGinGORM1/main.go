package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:haoyanroot@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 创建表，自动迁移
	//db.AutoMigrate(&UserInfo{})
	//u1 := UserInfo{1, "Tom", "男", "游泳"}
	//db.Create(u1)

	// 查询数据库表中第一条数据
	var u UserInfo
	db.First(&u)
	fmt.Println(u)

	// 更新
	db.Model(&u).Update("hobby", "双色球")

	// 删除
	db.Delete(&u)
}
