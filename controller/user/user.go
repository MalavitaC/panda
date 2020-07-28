package user

import (
	"log"
	"net/http"
	"panda/service"

	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "无",
	})
}

type LoginParams struct {
	Code string `json:"code"`
}

func Login(c *gin.Context) {

	var body LoginParams
	c.BindJSON(&body)

	wxUser, err := service.GetOpenID(body.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	log.Printf("%+v\n", wxUser)

	user := model.findOrCreateUserByOpenID(wxUser)
	// log.Println(user)

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"name":   "蔡文心",
	})
}
