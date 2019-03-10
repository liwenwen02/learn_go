package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*func GetDescription(c *gin.Context) {
	resp := map[string]string{"hello":"world"}
	c.JSON(http.StatusOK, resp)
	c.Next()
}*/
func GetDescription() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := map[string]string{"hello":"world"}
		ctx.JSON(http.StatusOK, resp)
		ctx.Next()
	}
}
func main() {
	/*r := gin.Default()
	r.GET("/description", GetDescription)
	r.Run(":8088")*/
	r := gin.New()
	r.Use(GetDescription())
	r.Run(":8087")
}