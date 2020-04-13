package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine  {
	gin.ForceConsoleColor()
	r:=gin.Default()
	
	r.GET("todo", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"test",
		})
	})
	return r
}
