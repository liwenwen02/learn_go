package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func MiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("---middle----")
		ctx.Set("request","client_request")
		ctx.Next()
		fmt.Println("---middle----")
	}
}
func main(){
	router := gin.Default()
	router.Use(MiddleWare())
	router.GET("/middleware", func(ctx *gin.Context) {
		request := ctx.MustGet("request").(string)
		req, _ := ctx.Get("request")
		ctx.JSON(http.StatusOK,gin.H{
			"middle_request":request,
			"request":req,
		})
		})
	router.Run("8097")
}
