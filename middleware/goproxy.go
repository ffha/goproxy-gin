package middleware

import (
	"github.com/goproxy/goproxy"
	"github.com/gin-gonic/gin"
)

func Goproxy() gin.HandlerFunc {
	return gin.WrapH(&goproxy.Goproxy{})
}