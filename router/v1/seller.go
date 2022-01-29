package v1

import (
	"coupon/app"
	"coupon/router/v1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SellerCreateHandler(c *gin.Context) {
	ctx := app.Wrap(c)
	request := service.CreateUserRequest{}

	err := ctx.BindParams(request)
	if err != nil {
		ctx.ToErrorResponse(http.StatusBadRequest, err)
		return
	}

	svc := service.New()
	err = svc.CreateSeller(request)
	if err != nil {
		ctx.ToErrorResponse(http.StatusInternalServerError, err)
		return
	}

	ctx.ToSuccessResponse(map[string]interface{}{
		"message": "create seller success",
	})

}
