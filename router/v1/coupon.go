package v1

import (
	"coupon/app"
	"coupon/router/v1/service"
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
	request.SellerID = uint32(c.GetUint("user_id"))

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
