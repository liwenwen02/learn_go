package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", func(ctx *gin.Context) {
		firstname := ctx.DefaultQuery("firstname","Guest")
		lastname := ctx.Query("lastname")
		ctx.String(http.StatusOK, "%s %s",firstname, lastname)
	})
	router.Run(":8080")
}
