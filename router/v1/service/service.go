package service

import (
	"coupon/common"
	"coupon/data"
	"coupon/model"
	"time"

	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func New() *service {
	return &service{
		db: data.DB,
	}
}

func (s *service) CreateUser(req CreateUserRequest) error {
	user := model.User{
		UserName: req.UserName,
		Password: common.MD5(req.Password),
	}
	return user.Create(s.db)
}

func (s *service) CreateSeller(req CreateUserRequest) error {
	seller := model.Seller{
		UserName: req.UserName,
		Password: common.MD5(req.Password),
	}
	return seller.Create(s.db)
}

func (s *service) UserLogin(req LoginRequest) (uint32, error) {
	user := model.User{
		UserName: req.UserName,
		Password: common.MD5(req.Password),
	}
	err := user.Validate(s.db)
	if err != nil {
		return 0, err
	}

	return user.ID, err
}

func (s *service) SellerLogin(req LoginRequest) (uint32, error) {
	seller := model.Seller{
		UserName: req.UserName,
		Password: common.MD5(req.Password),
	}
	err := seller.Validate(s.db)
	return seller.ID, err
}

func (s *service) CreateCoupon(req CreateCouponRequest) error {
	coupon := model.Coupon{
		Name:      req.CouponName,
		Amount:    *req.Amount,
		Left:      *req.Amount,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(30 * 24 * time.Hour),
		SellerID:  req.SellerID,
	}
	return coupon.Create(s.db)
}
