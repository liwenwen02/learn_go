package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/form_post", func(ctx *gin.Context) {
		message := ctx.PostForm("message")
		ctx.JSON(http.StatusOK,gin.H{
			"status":gin.H{
				"status_code": http.StatusOK,
				"status":"ok",
			},
			"message": message,
		})
	})
	router.Run(":8082")
}
