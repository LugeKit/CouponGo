package service

type CreateUserRequest struct {
	UserName string `form:"user_name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginRequest struct {
	UserName string `form:"user_name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type CreateCouponRequest struct {
	CouponName string `form:"name" binding:"required"`
	Amount     *int   `form:"amount" binding:"required"`
	SellerID   uint32
}
