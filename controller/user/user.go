package user

import (
	"log"
	"net/http"
	"panda/model"
	"panda/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	Id     uint64 `json:"id"`
	OpenID string `json: "openID"`
}

func Query(c *gin.Context) {
	var body QueryParams
	var result *model.User
	c.BindQuery(&body)

	if body.Id == 0 && body.OpenID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}

	if body.Id > 0 {
		result = model.QueryUserById(body.Id)
	} else {
		result = model.QueryUserByOpenID(body.OpenID)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"data":   result,
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
			"message": err.Error(),
		})
		return
	}

	user := model.FindOrCreateUserByOpenID(wxUser)
	log.Println(user)
	needSync := false
	if user.NickName == "" {
		needSync = true
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
	Gender    int8   `json:"gender"`
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
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "openID is not found ",
		})
		return
	}

	model.UpdateByOpenID(model.User{
		NickName:  body.NickName,
		Gender:    body.Gender,
		Language:  body.Language,
		City:      body.City,
		Province:  body.Province,
		Country:   body.Country,
		AvatarUrl: body.AvatarUrl,
	}, body.OpenID)

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
	})
}

type DeleteParams struct {
	Id     uint64 `json:"id"`
	OpenID string `json: "openID"`
}

func Delete(c *gin.Context) {
	var body DeleteParams
	c.BindJSON(&body)

	if body.Id == 0 && body.OpenID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}

	if body.Id > 0 {
		model.DeleteUserById(body.Id)
	} else {
		model.DeleteUserByOpenID(body.OpenID)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"data":   body,
	})
}

type UpdateParams struct {
	NickName  string `json:"nickName"`
	Gender    int8   `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
}

func Update(c *gin.Context) {
	var body UpdateParams
	c.BindJSON(&body)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 0)

	if user := model.QueryUserById(id); user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不存在",
		})
		return
	}
	model.UpdateById(model.User{
		NickName:  body.NickName,
		Gender:    body.Gender,
		Language:  body.Language,
		City:      body.City,
		Province:  body.Province,
		Country:   body.Country,
		AvatarUrl: body.AvatarUrl,
	}, id)

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"data":   body,
	})
}
