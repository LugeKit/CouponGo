package v1

import (
	"coupon/app"
	"coupon/common"
	"coupon/conf"
	"coupon/router/v1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserCreateHandler(c *gin.Context) {
	ctx := app.Wrap(c)
	request := service.CreateUserRequest{}

	err := ctx.BindParams(&request)
	if err != nil {
		ctx.ToErrorResponse(http.StatusBadRequest, err)
		return
	}

	svc := service.New()
	err = svc.CreateUser(request)
	if err != nil {
		ctx.ToErrorResponse(http.StatusInternalServerError, err)
		return
	}

	ctx.ToSuccessResponse(map[string]interface{}{
		"message": "create user success",
	})
}

func UserLoginHandler(c *gin.Context) {
	ctx := app.Wrap(c)
	request := service.LoginRequest{}

	err := ctx.BindParams(&request)
	if err != nil {
		ctx.ToErrorResponse(http.StatusBadRequest, err)
		return
	}

	svc := service.New()
	userID, err := svc.UserLogin(request)
	if err != nil {
		ctx.ToErrorResponse(http.StatusInternalServerError, err)
		return
	}

	token, err := common.GenerateToken(conf.AppConfig.JWT.UserSecret, userID)
	c.Header("Authorization", token)
	ctx.ToSuccessResponse(map[string]interface{}{
		"message": "login success",
	})
}
