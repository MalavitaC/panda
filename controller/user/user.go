package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "无",
	})
}

type login struct {
	Code string `json:"code"`
}

func Login(c *gin.Context) {

	var login login

	c.BindJSON(&login)

	log.Println(login.Code)
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"name":   "蔡文心",
	})
}
