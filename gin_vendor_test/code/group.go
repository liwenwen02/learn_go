package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/login", func(ctx *gin.Context) {
		ctx.String(http.StatusOK,"v1 login")
	})

	v2 := router.Group("/v2")
	v2.GET("/login", func(ctx *gin.Context) {
		ctx.String(http.StatusOK,"v2 login")
	})
	router.Run(":8080")
}
