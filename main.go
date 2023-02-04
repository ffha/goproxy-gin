package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goproxy/goproxy"
)

func main(){
	r := gin.New()
	r.Use(gin.Logger())
    r.Use(gin.Recovery())
	r.Use(gin.WrapH(&goproxy.Goproxy{}))
	r.Run()
}