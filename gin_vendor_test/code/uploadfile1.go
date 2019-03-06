package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../templates/*")
	router.GET("/upload", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK,"upload.html", gin.H{})
	})
	router.Run(":8017")
}
