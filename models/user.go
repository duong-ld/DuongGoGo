package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Base
	BirthDay time.Time
	Email    string `gorm:"unique"`
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	zeroTime, err := time.Parse("2006-01-02", "0000-00-00")

	if u.BirthDay == zeroTime {
		u.BirthDay, err = time.Parse("2006-01-02", "1970-01-01")
	}

	return
}
