package router

import (
	"gin-todo/model"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine  {
	gin.ForceConsoleColor()
	r:=gin.Default()
	model.CreateTodo(model.Todo{"TestId","TestDescription"})
	r.GET("todo", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"test",
		})
	})
	return r
}
