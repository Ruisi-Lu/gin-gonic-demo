package router

import (
	"gin-gonic-demo/handlers"
	"github.com/gin-gonic/gin"
)

func Init() {

	r := gin.Default()
	// 設定群組 /demo   eg. http://0.0.0.0/demo/
	demo := r.Group("/demo")
		// 子目錄 eg. http://0.0.0.0/demo/123
		demo.GET("/", handlers.Welcome)
		demo.GET("123", handlers.HelloPage)
		demo.GET("a", handlers.HelloPage2)
		// 參數傳遞 (*或:)
		// 必給參數 不給會404
		demo.GET("/hello/:name", handlers.HelloPage3)
		// 可不給參數
		demo.GET("/hello2/*name", handlers.HelloPage3)
		// post demo
		demo.POST("/posthello/:name", handlers.HelloPage4)

	// 中間件 BasicAuth
	account := r.Group("/account", gin.BasicAuth(gin.Accounts{
		// 測試帳號
		"ruisi":"123",
		"wayne":"234",
	}))
	
	account.GET("", handlers.Login)
	account.GET(":name", handlers.Loginwelcome)
	r.Run(":80")

}