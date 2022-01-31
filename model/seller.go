package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Seller struct {
	ID       uint32 `gorm:"id;primary_key;auto_increment"`
	UserName string `gorm:"user_name;type:varchar(20);unique"`
	Password string `gorm:"password;type:varchar(32)"`
}

func (s Seller) Create(db *gorm.DB) error {
	return db.Create(&s).Error
}

func (s *Seller) Validate(db *gorm.DB) error {
	err := db.Where("user_name = ? and password = ?", s.UserName, s.Password).First(s).Error
	if err != nil {
		return err
	}

	if s.ID == 0 {
		return fmt.Errorf("username or password not right!")
	}

	return nil
}
