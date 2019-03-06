package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/user/:name", func(ctx *gin.Context){
		name := ctx.Param("name")
		ctx.String(http.StatusOK,"welcome %s", name)
	})

	router.GET("/user/:name/*action",func(ctx *gin.Context){
		name := ctx.Param("name")
		action := ctx.Param("action")
		ctx.String(http.StatusOK, "%s", name + " " + action)
	})
	router.Run(":8080")
}
