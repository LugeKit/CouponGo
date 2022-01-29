package model

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	ID        uint32    `gorm:"id;primary_key;auto_increment;"`
	Name      string    `gorm:"name;type:varchar(60)"`
	Amount    int       `gorm:"amount;not null"`
	Left      int       `gorm:"left; not null"`
	CreatedAt time.Time `gorm:"created_at"`
	ExpiredAt time.Time `gorm:"expired_at"`
	SellerID  uint32    `gorm:"seller_id"`
	Seller    Seller
}

func (c *Coupon) Create(db *gorm.DB) error {
	return db.Create(c).Error
}
