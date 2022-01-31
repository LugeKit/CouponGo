package middleware

import (
	"coupon/app"
	"coupon/common"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SellerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := app.Wrap(c)
		token := c.GetHeader("Authorization")
		log.Println(token)
		claims, err := common.ParseToken(token)
		if err != nil {
			ctx.ToErrorResponse(http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
	}
}
