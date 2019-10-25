package router

import (
	"../handlers"
	"github.com/gin-gonic/gin"
)

func Init() {

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("123", handlers.HelloPage)
		v1.GET("a", handlers.HelloPage2)
		v1.GET("/hello/:name", handlers.HelloPage3)
		v1.POST("/hello/:name", handlers.HelloPage4)
	}
	r.Run(":80")

}
