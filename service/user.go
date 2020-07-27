package service

import "panda/model"

func QueryUserByOpenID(OpenID string) {
	var user model.User

	model.DB.Where("openID = ?", OpenID).Find()
}
