package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main(){
	router := gin.Default()
	router.POST("/upload", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
		}
	})
}
