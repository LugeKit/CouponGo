package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppContext struct {
	c *gin.Context
}

func Wrap(c *gin.Context) AppContext {
	return AppContext{
		c: c,
	}
}

func (ctx *AppContext) BindParams(requestParams interface{}) error {
	if err := ctx.c.Bind(requestParams); err != nil {
		return err
	}
	return nil
}

func (ctx *AppContext) ToSuccessResponse(result interface{}) {
	ctx.c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (ctx *AppContext) ToErrorResponse(errcode int, err error) {
	ctx.c.JSON(errcode, gin.H{
		"status": "fail",
		"error":  err.Error(),
	})
}
