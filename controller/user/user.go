package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "无",
	})
}

func Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "用户已创建",
	})
}
