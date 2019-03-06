package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		name := ctx.Query("name")
		age := ctx.DefaultQuery("age","10")
		bodyinfo := ctx.PostForm("info")
		ctx.JSON(http.StatusOK, gin.H{
			"name":name,
			"age": age,
			"bodyinfo": bodyinfo,
		})

	})

	router.Run(":8080")
}
