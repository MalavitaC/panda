package main

import (
	"fmt"
	"panda/model"
	"panda/router"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	app := gin.Default()
	router.Register(app)

	model.InitDB()

	app.Run("0.0.0.0:9000")
	fmt.Println("端口:", 9000)

}
