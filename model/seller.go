package model

import "gorm.io/gorm"

type Seller struct {
	ID       uint32 `gorm:"id;primary_key;auto_increment"`
	UserName string `gorm:"user_name;type:varchar(20);unique"`
	Password string `gorm:"password;type:varchar(32)"`
}

func (s Seller) Create(db *gorm.DB) error {
	return db.Create(&s).Error
}
