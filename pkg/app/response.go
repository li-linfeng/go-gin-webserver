package app

import (
	"meta/pkg/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	HttpCode  int
	ErrorCode int
	Message   string
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Respond() *Gin {
	return &Gin{200, 200, ""}
}

func (g *Gin) WithHttpCode(code int) *Gin {
	g.HttpCode = code
	return g
}

func (g *Gin) WithErrorCode(code int) *Gin {
	g.ErrorCode = code
	return g
}

func (g *Gin) WithErrorMessage(message string) *Gin {
	g.Message = message
	return g
}

func (g *Gin) Success(c *gin.Context, data interface{}) {
	c.JSON(g.HttpCode, Response{
		Code: 200,
		Msg:  e.GetMsg(200),
		Data: data,
	})
}

func (g *Gin) Fail(c *gin.Context, data interface{}) {
	c.JSON(g.HttpCode, Response{
		Code: 400,
		Msg:  e.GetMsg(400),
		Data: data,
	})
}

func (g *Gin) Response(c *gin.Context, data interface{}) {
	msg := e.GetMsg(g.ErrorCode)
	if len(g.Message) > 0 {
		msg = g.Message
	}
	c.JSON(g.HttpCode, Response{
		Code: g.ErrorCode,
		Msg:  msg,
		Data: data,
	})
}
