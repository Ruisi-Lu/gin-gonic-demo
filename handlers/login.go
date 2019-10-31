package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
// Login 簡單BasicAuth
func Login(context *gin.Context) {
	context.String(http.StatusOK, "welcome login demo")
}