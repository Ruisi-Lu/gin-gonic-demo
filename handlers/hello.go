package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Welcome 歡迎訊息
// @Summary 歡迎訊息
// @Tags Hello API
// @version 1.0
// @Accept json
// @Success 200 "welcome demo"
// @Router /demo/ [get]
func Welcome(context *gin.Context) {
	context.String(http.StatusOK, "welcome demo")
}

// HelloPage 歡迎訊息Json
// @Summary 歡迎訊息Json
// @Tags Hello API
// @version 1.0
// @Accept json
// @Produce  json
// @Success 200 "{"message":"welcome"}"
// @Router /demo/123 [get]
func HelloPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "welcome",
	})
}

// HelloPage3 輸入值回傳(限定)
// @Summary 輸入值回傳(限定)
// @Tags Hello API
// @verion 1.0
// @Accept json
// @Param name path string true "name"
// @Success 200 "Hello name"
// @Router /demo/hello/{name} [get]
func HelloPage3(context *gin.Context) {
	name := context.Param("name")
	context.String(http.StatusOK, "Hello %s", name)
}
// HelloPage3_1 輸入值回傳(不限定)
// @Summary 輸入值回傳(不限定)
// @Tags Hello API
// @verion 1.0
// @Accept json
// @Param name path string false "name"
// @Success 200 "Hello name"
// @Router /demo/hello2/{name} [get]
func HelloPage3_1(context *gin.Context) {
	name := context.Param("name")
	context.String(http.StatusOK, "Hello %s", name)
}

// HelloPage4 輸入值回傳(限定)
// @Summary 輸入值回傳(限定)
// @Tags Hello API
// @verion 1.0
// @Accept json
// @Param name path string true "name"
// @Success 200 "Hello name"
// @Router /demo/posthello/{name} [post]
func HelloPage4(context *gin.Context) {
	name := context.Param("name")
	context.String(http.StatusOK, "Hello %s", name)
}
