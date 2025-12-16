package main

import (
	"crud/internal/controller"
	"crud/internal/midware"
	"crud/pkg/db"
	"crud/pkg/response"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	//db.DB.AutoMigrate(&model.User{}) //自动建表

	r := gin.Default()
	//全局日志
	r.Use(midware.Logger())
	//公开接口
	controller.RegisterUserRoutes(r)

	//需要登录的接口
	//1.创建需要认证的路由组
	auth := r.Group("/api")
	//2.为该路由组添加jwt认证中间件
	auth.Use(midware.JWTAuth())

	//3.在该路由组下定义需要认证的接口
	{
		auth.GET("/profile", func(c *gin.Context) {
			uid, _ := c.Get("user_id")
			response.Success(c, gin.H{"user_id": uid})
		})
	}

	r.Run(":8080")
}
