package model

import (
	"panda/service"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID         uint                   `gorm:"cloumn:id;primary_key"`
	OpenID     string                 `gorm:"cloumn:openID;size:256;NOT NULL"`
	UnionID    string                 `gorm:"cloumn:unionID;size:256;NOT NULL;DEFAULT:\"\""`
	SessionKey string                 `gorm:"cloumn:sessionKey;size:256;NOT NULL"`
	Name       string                 `gorm:"cloumn:namea;size:100;NOT NULL;DEFAULT:\"\""`
	NickNamea  string                 `gorm:"cloumn:nickNamea;size:100;NOT NULL;DEFAULT:\"\""`
	Vatarurl   string                 `gorm:"cloumn:vatarurl;size:100;NOT NULL;DEFAULT:\"\""`
	Country    string                 `gorm:"cloumn:country;size:100;NOT NULL;DEFAULT:\"\""`
	Province   string                 `gorm:"cloumn:province;size:100;NOT NULL;DEFAULT:\"\""`
	City       string                 `gorm:"cloumn:city;size:100;NOT NULL;DEFAULT:\"\""`
	Language   string                 `gorm:"cloumn:language;size:100;NOT NULL;DEFAULT:\"\""`
	Mobile     string                 `gorm:"cloumn:mobile;size:50;NOT NULL;DEFAULT:\"\""`
	Telnum     string                 `gorm:"cloumn:telnum;size:13;NOT NULL;DEFAULT:\"\""`
	Exif       map[string]interface{} `gorm:"cloumn:exif;type:json;DEFAULT:null"`
	Status     uint                   `gorm:"cloumn:status;NOT NULL;DEFAULT:1"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}

func findOrCreateUserByOpenID(wxUser service.BodyStruct) *gorm.DB | nil {
	var user User

	result := DB.Where(User{OpenID: wxUser.Openid}).Attrs(User{SessionKey: wxUser.Session_key}).FirstOrCreate(&user)
	if gorm.IsRecordNotFoundError(result.Error) {
		return nil
	}
	return result
}
