package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context) {
	context.String(http.StatusOK, "welcome login demo")
}

func Loginwelcome(context *gin.Context) {
	name := context.Param("name")
	context.String(http.StatusOK, "welcome login demo %s", name)
}