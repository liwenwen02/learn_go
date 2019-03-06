package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
	"net/http"
)

func main() {
	router := gin.Default()
	var u user
	router.POST("/", func(ctx *gin.Context) {
		err := ctx.Bind(&u)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"name": u.name,
			"age": u.age,
		})
})
	router.Run()
}
