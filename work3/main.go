package main

import (
	"github.com/gin-gonic/gin"
	"simpleserver/work3/handler"
	"simpleserver/work3/middleware"
)

func main() {
	router := gin.Default()
	router.Use(middleware.IsCached)
	router.GET("/price/:currency", handler.GetMessage)
	router.Run("0.0.0.0:8080")
}
