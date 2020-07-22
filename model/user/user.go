package User

import "time"

type User struct {
	ID        uint                   `gorm:"cloumn:id;primary_key"`
	OpenID    string                 `gorm:"cloumn:openID;size:28;NOT NULL"`
	NickNamea string                 `gorm:"cloumn:nickNamea;size:100;NOT NULL"`
	Vatarurl  string                 `gorm:"cloumn:vatarurl;size:100;NOT NULL"`
	Country   string                 `gorm:"cloumn:country;size:100;NOT NULL"`
	Province  string                 `gorm:"cloumn:province;size:100;NOT NULL"`
	City      string                 `gorm:"cloumn:city;size:100;NOT NULL"`
	Language  string                 `gorm:"cloumn:language;size:100;NOT NULL"`
	Mobile    string                 `gorm:"cloumn:mobile;size:50;NOT NULL"`
	Telnum    string                 `gorm:"cloumn:telnum;size:13;NOT NULL"`
	Exif      map[string]interface{} `gorm:"cloumn:exif;type:json;DEFAULT:null"`
	Status    uint                   `gorm:"cloumn:status;NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
