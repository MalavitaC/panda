package main

import (
	"fmt"
	"panda/router"
	"panda/model/user/User"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	defer func ()  {
		
	}
	app := gin.Default()
	router.Register(app)

	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		panic("数据库连接失败")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	app.Run("0.0.0.0:9000")
	fmt.Println("端口:", 9000)

}
