package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {

	db, err := gorm.Open("mysql", "root:73056qaz@(106.14.60.154:8306)/panda?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")

	// defer db.Close()

	// Migrate the schema
	db.AutoMigrate()

	DB = db
}
