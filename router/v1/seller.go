package v1

import (
	"coupon/app"
	"coupon/common"
	"coupon/conf"
	"coupon/router/v1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SellerCreateHandler(c *gin.Context) {
	ctx := app.Wrap(c)
	createReq := service.CreateUserRequest{}

	err := ctx.BindParams(&createReq)
	if err != nil {
		ctx.ToErrorResponse(http.StatusBadRequest, err)
		return
	}

	svc := service.New()
	err = svc.CreateSeller(createReq)
	if err != nil {
		ctx.ToErrorResponse(http.StatusInternalServerError, err)
		return
	}

	ctx.ToSuccessResponse(map[string]interface{}{
		"message": "create seller success",
	})

}

func SellerLoginHandler(c *gin.Context) {
	ctx := app.Wrap(c)
	loginReq := service.LoginRequest{}

	err := ctx.BindParams(&loginReq)
	if err != nil {
		ctx.ToErrorResponse(http.StatusBadRequest, err)
		return
	}

	svc := service.New()
	sellerID, err := svc.UserLogin(loginReq)
	if err != nil {
		ctx.ToErrorResponse(http.StatusInternalServerError, err)
		return
	}

	token, err := common.GenerateToken(conf.AppConfig.JWT.SellerSecret, sellerID)
	if err != nil {
		ctx.ToErrorResponse(http.StatusInternalServerError, err)
		return
	}
	c.Header("Authorization", token)
	ctx.ToSuccessResponse(map[string]interface{}{
		"message": "login success",
	})
}
