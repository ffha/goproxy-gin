package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ffha/goproxy-gin/middleware"
)


func main(){
	r := gin.New()
	r.TrustedPlatform = "X-Forwarded-For"
	r.Use(gin.Logger())
    r.Use(gin.Recovery())
	r.Use(middleware.Goproxy())
	r.Use(middleware.Cache())
	r.Run()
}