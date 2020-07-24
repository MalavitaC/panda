package main

import (
	"fmt"
	user "panda/model/user"
	"panda/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	app := gin.Default()
	router.Register(app)

	db, err := gorm.Open("mysql", "root:73056qaz@(106.14.60.154:8306)/panda?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("数据库连接失败")
	}
	fmt.Println("数据库连接成功")

	// defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&user.User{})

	app.Run("0.0.0.0:9000")
	fmt.Println("端口:", 9000)

}
