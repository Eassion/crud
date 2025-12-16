package controller

import (
	"crud/internal/dto"
	"crud/internal/service"
	"crud/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	user := router.Group("/users")
	{
		user.POST("", CreateUser)
		user.GET("/:id", GetUser)
		user.GET("", GetUserList)
		user.PUT("/:id", UpdateUser)
		user.DELETE("/:id", DeleteUser)
	}
}

func DeleteUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))

	err := service.DeleteUser(uint(id))
	if err != nil {
		response.Fail(context, err.Error())
		return
	}
	response.Success(context, nil)
}

func UpdateUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))

	var req dto.UpdateUserRequest

	_ = context.ShouldBindJSON(&req)

	err := service.UpdateUser(uint(id), req.Age)
	if err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
}

func GetUserList(context *gin.Context) {
	users, _ := service.GetUserList()
	response.Success(context, users)
}

func GetUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	user, err := service.GetUser(uint(id))
	if err != nil {
		response.Fail(context, "用户不存在")
		return
	}
	response.Success(context, user)
}

func CreateUser(context *gin.Context) {
	var req dto.CreateUserRequest

	if err := context.ShouldBindJSON(&req); err != nil { //  ;分割变量初始化和条件判断  作用：变量的作用域只在if语句中，不会污染外部作用域
		response.Fail(context, err.Error())
		return
	}

	err := service.CreateUser(req.Username, req.Password, req.Age)
	if err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
}
