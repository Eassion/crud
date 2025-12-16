package main

import (
	"crud/internal/controller"
	"crud/internal/model"
	"crud/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	db.DB.AutoMigrate(&model.User{}) //自动建表

	r := gin.Default()

	controller.RegisterUserRoutes(r)

	r.Run(":8080")
}
