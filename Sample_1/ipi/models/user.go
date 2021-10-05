package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          int64  `gorm:"type:integer;not null:primarykey"`
	Fullname    string `gorm :"type:varchar(255) ;not null"`
	Username    string `gorm :"type:varchar(255) ;not null"`
	Gender      string `gorm :"type:varchar(10) ;not null"`
	Birthday    string `gorm :"type:varchar(10) ;not null"`
	CreatedAt   time.Time
	UpdateAt    time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Auth_tokens []*Auth_token  `gorm:"foreignKey:UserID;references:Id"`
}

type Auth_token struct {
	Id        int64  `gorm:"type:integer;not null:primarykey"`
	UserID    int64  `gorm:"type:integer;not null"`
	Tocken    string `gorm :"type:varchar(255) ;not null"`
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// type Contact struct {
// 	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()`
// 	PhoneNumber string    `gorm:"type:varchar(12)"`
// 	Email       string    `gorm:"type:varchar(256)"`
// 	Fax         string    `gorm:"type:varchar(256)"`
// 	PeopleID    uuid.UUID `gorm:"type:uuid;not null"`
// 	CreatedAt   time.Time
// 	UpdatedAt   time.Time
// 	DeletedAt   gorm.DeletedAt `gorm:"index"`
// }

// type People struct {
// 	Id             uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4()`
// 	Name           string          `gorm:"type:varchar(256)"`
// 	Slut           string          `gorm:"type:varchar(256);not null;unquie"`
// 	Age            int64           `gorm:"type:integer"`
// 	Address        sql.NullString  `gorm:"type:varchar(256)"`
// 	AccountBalance sql.NullFloat64 `gorm:"type:real"`
// 	Contacts       []*Contact      `gorm:"foreignKey:PeopleID;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` //Delete luôn Contacts nếu People bị Delete
// 	CreatedAt      time.Time
// 	Updat
