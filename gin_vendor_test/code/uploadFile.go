package main

import (
	"github.com/gin-gonic/gin"
	"log"
		"net/http"
	"os"
	"io"
	"fmt"
)

func main() {
	router := gin.Default()
	router.POST("/uploadfile", func(ctx *gin.Context) {
		file, header, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.String(http.StatusOK,"upload err")
			log.Fatal(err)
			return
		}
		filename := header.Filename
		fmt.Print("received file:", filename)
		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		ctx.String(http.StatusOK,"uploadsuccessful")
	})
	router.Run(":8099")
}