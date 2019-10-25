package router

import (
	"gin-gonic-demo/handlers"
	"github.com/gin-gonic/gin"
)

func Init() {

	r := gin.Default()
	// 設定群組 /demo   eg. http://0.0.0.0/demo/
	demo := r.Group("/demo")
	{
		// 子目錄 eg. http://0.0.0.0/demo/123
		demo.GET("/", handlers.Welcome)
		demo.GET("123", handlers.HelloPage)
		demo.GET("a", handlers.HelloPage2)
		demo.GET("/hello/:name", handlers.HelloPage3)
		demo.POST("/hello/:name", handlers.HelloPage4)
	}
	r.Run(":80")

}