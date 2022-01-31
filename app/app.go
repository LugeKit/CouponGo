package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	c *gin.Context
}

func Wrap(c *gin.Context) Context {
	return Context{
		c: c,
	}
}

func (ctx *Context) BindParams(requestParams interface{}) error {
	if err := ctx.c.Bind(requestParams); err != nil {
		return err
	}
	return nil
}

func (ctx *Context) ToSuccessResponse(result interface{}) {
	ctx.c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (ctx *Context) ToErrorResponse(errcode int, err error) {
	ctx.c.JSON(errcode, gin.H{
		"status": "fail",
		"error":  err.Error(),
	})
}
