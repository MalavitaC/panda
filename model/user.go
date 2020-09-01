package model

import (
	"panda/service"
	"time"
)

type User struct {
	ID         uint64                 `gorm:"Column:id;primary_key"`
	OpenID     string                 `gorm:"Column:openID;size:256;NOT NULL"`
	UnionID    string                 `gorm:"Column:unionID;size:256;NOT NULL;DEFAULT:\"\""`
	SessionKey string                 `gorm:"Column:sessionKey;size:256;NOT NULL"`
	Name       string                 `gorm:"Column:namea;size:100;NOT NULL;DEFAULT:\"\""`
	NickName   string                 `gorm:"Column:nickName;size:100;NOT NULL;DEFAULT:\"\""`
	Gender     int8                   `gorm:"Column:gender;NOT NULL;"`
	AvatarUrl  string                 `gorm:"Column:avatarurl;size:256;NOT NULL;DEFAULT:\"\""`
	Country    string                 `gorm:"Column:country;size:100;NOT NULL;DEFAULT:\"\""`
	Province   string                 `gorm:"Column:province;size:100;NOT NULL;DEFAULT:\"\""`
	City       string                 `gorm:"Column:city;size:100;NOT NULL;DEFAULT:\"\""`
	Language   string                 `gorm:"Column:language;size:100;NOT NULL;DEFAULT:\"\""`
	Mobile     string                 `gorm:"Column:mobile;size:50;NOT NULL;DEFAULT:\"\""`
	Telnum     string                 `gorm:"Column:telnum;size:13;NOT NULL;DEFAULT:\"\""`
	Exif       map[string]interface{} `gorm:"Column:exif;type:json;DEFAULT:null"`
	Status     uint                   `gorm:"Column:status;NOT NULL;DEFAULT:1"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}

func FindOrCreateUserByOpenID(wxUser service.BodyStruct) User {
	var user User
	DB.Where(User{OpenID: wxUser.Openid}).Attrs(User{SessionKey: wxUser.Session_key}).FirstOrCreate(&user)
	return user
}

func QueryUserById(Id uint64) *User {
	var user User
	if DB.Where(User{ID: Id}).First(&user).RecordNotFound() {
		return nil
	}
	return &user
}

func QueryUserByOpenID(openID string) *User {
	var user User
	if DB.Where(User{OpenID: openID}).First(&user).RecordNotFound() {
		return nil
	}
	return &user
}

func UpdateByOpenID(params User, openID string) {
	var user User
	DB.Model(&user).Where("openID = ?", openID).Updates(params)
}

func DeleteUserById(Id uint64) {
	DB.Where(User{ID: Id}).Delete(&User{})
}

func DeleteUserByOpenID(openID string) {
	DB.Where(User{OpenID: openID}).Delete(&User{})
}

func UpdateById(params User, Id uint64) {
	var user User
	DB.Model(&user).Where("id = ?", Id).Updates(params)
}
