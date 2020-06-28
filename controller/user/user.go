package user

import "github.com/gin-gonic/gin"

func Query(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
