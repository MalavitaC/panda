package main

import (
	"panda/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	router.Register(app)
	app.Run()
}
