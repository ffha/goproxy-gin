package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cache"
	"time"
)

func Cache() gin.HandlerFunc{
    store := persistence.NewInMemoryStore(time.Second)
    return cache.CachePage(store, time.Minute)
}