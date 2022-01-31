package v1

import (
	"coupon/app"
	"coupon/router/v1/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CouponCreateHandler(c *gin.Context) {
	ctx := app.Wrap(c)
	request := service.CreateCouponRequest{}

	err := ctx.BindParams(&request)
	if err != nil {
		ctx.ToErrorResponse(http.StatusBadRequest, err)
		return
	}
	sellerID, ok := c.Get("user_id")
	if !ok {
		ctx.ToErrorResponse(http.StatusUnauthorized, fmt.Errorf("can't get token"))
		return
	}

	request.SellerID, ok = sellerID.(uint32)
	if !ok {
		ctx.ToErrorResponse(http.StatusInternalServerError, fmt.Errorf("token wrong"))
		return
	}

	svc := service.New()
	err = svc.CreateCoupon(request)
	if err != nil {
		ctx.ToErrorResponse(http.StatusInternalServerError, err)
		return
	}

	ctx.ToSuccessResponse(map[string]interface{}{
		"message": "create coupon success!",
	})
}
