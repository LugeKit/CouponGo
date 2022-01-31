package router

import (
	"coupon/middleware"
	v1 "coupon/router/v1"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	apiv1 := server.Group("/apiv1")
	{
		apiv1.Handle("POST", "/user", v1.UserCreateHandler)
		apiv1.Handle("POST", "/seller", v1.SellerCreateHandler)
		apiv1.Handle("POST", "/login/user", v1.UserLoginHandler)
		apiv1.Handle("POST", "/login/seller", v1.SellerLoginHandler)

		couponCreate := apiv1.Group("/seller")
		couponCreate.Use(middleware.SellerAuth())
		couponCreate.Handle("POST", "/coupon", v1.CouponCreateHandler)
	}

	return server
}
