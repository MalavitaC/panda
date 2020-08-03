package user

import (
	"log"
	"net/http"
	"panda/model"
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

	user := model.FindOrCreateUserByOpenID(wxUser)
	log.Println(user)
	needSync := true
	if user.NickName == "" {
		needSync = false
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "SUCCESS",
		"openID":   user.OpenID,
		"needSync": needSync,
	})
}

type SyncUserInfoParams struct {
	OpenID    string `json:"openID"`
	NickName  string `json:"nickName"`
	Gender    int32  `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
}

func SyncUserInfo(c *gin.Context) {

	var body SyncUserInfoParams
	c.BindJSON(&body)

	user := model.QueryUserByOpenID(body.OpenID)
	log.Printf("%+v\n", user)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "openID is not found ",
		})
		return
	}

	model.UpdateByOpenID(body, body.OpenID)

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
	})
}
