package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simpleserver/work3/cache"
)

// 中间件的方式实现缓存逻辑
func IsCached(c *gin.Context) {
	key := c.Param("currency")
	value, ok := cache.DefaultCache.Get(key)
	fmt.Println("进入中间件")
	if ok {
		c.Data(http.StatusOK, "application/json", []byte(value))
		fmt.Println("缓存已存在")
		c.Abort()
		return
	}
	c.Next()
	fmt.Println("设置缓存")
	v, _ := c.Get("content")
	value = string(v.([]byte))
	cache.DefaultCache.Set(key, value)
	fmt.Printf("Cache key %v\n", key)
}
