package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	ID       uint32 `gorm:"id;primary_key;auto_increment"`
	UserName string `gorm:"user_name;type:varchar(20);unique"`
	Password string `gorm:"password;type:varchar(32)"`
}

func (u User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u *User) Validate(db *gorm.DB) error {
	err := db.Where("user_name = ? and password = ?", u.UserName, u.Password).Find(u).Error
	if err != nil {
		return err
	}

	if u.ID == 0 {
		return fmt.Errorf("username or password not right!")
	}

	return nil
}
