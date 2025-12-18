package controller

import (
	"crud/internal/dto"
	"crud/internal/service"
	"crud/pkg/response"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	r.POST("/login", Login)
}

func Login(context *gin.Context) {
	var req dto.LoginRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		response.Fail(context, err.Error())
		return
	}

	token, err := service.Login(req.Username, req.Password)
	if err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, gin.H{"token": token})
}
