package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// type Context struct {
// 	*gin.Context
// }

// func getResponse(message string, data interface{}) gin.H {
// 	return gin.H{
// 		"success":  true,
// 		"message": message,
// 		"data":    data,
// 	}
// }

// func (c *Context) Response(status int, message string, data interface{}) {
// 	c.JSON(status, getResponse(message, data))
// }

type Response struct {
	Ctx *gin.Context
}

// TODO: refactor getResponse function
func getResponse(success bool, message string, data interface{}) gin.H {
	if success {
		return gin.H{
			"success": success,
			"message": message,
			"data":    data,
		}
	} else {
		return gin.H{
			"success": success,
			"message": message,
			"error":   data,
		}
	}
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) Success(message string, data interface{}) {
	r.Ctx.JSON(http.StatusOK, getResponse(true, message, data))
}

func (r *Response) Fail(message string, data interface{}) {
	r.Ctx.JSON(http.StatusInternalServerError, getResponse(false, message, data))
}
