package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func Welcome(context *gin.Context) {
	context.String(http.StatusOK, "welcome demo")
}
func HelloPage(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{
        "message": "welcome",
    })
}
func HelloPage2(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{
        "message": "123",
    })
}
func HelloPage3(context *gin.Context) {
	name := context.Param("name")
	context.String(http.StatusOK, "Hello %s", name)
}
func HelloPage4(context *gin.Context) {
	name := context.Param("name")
	context.String(http.StatusOK, "Hello %s", name)
}