package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	name string `json:name`
	age int `json:age`
}
func main () {
	router := gin.Default()
	var u user
	var err error
	router.POST("/", func(ctx *gin.Context) {
		content_type := ctx.Request.Header.Get("Content-Type")
		switch content_type {
			case "application/json":
				err = ctx.BindJSON(&u)
		    case "application/x-www-form-urlencoded":
		    	err = ctx.BindWith(&u, binding.Form)

		}
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"name": u.name,
			"age": u.age,
		})
	})
	router.Run(":8077")
}
