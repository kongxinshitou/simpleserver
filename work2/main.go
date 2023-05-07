package main

import (
	"github.com/gin-gonic/gin"
	"simpleserver/work3/handler"
)

func main() {
	router := gin.Default()
	router.GET("/price/:currency", handler.GetMessage)
	router.Run("0.0.0.0:8079")
}
