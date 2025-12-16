package dao

import (
	"crud/internal/model"
	"crud/pkg/db"
)

func CreateUser(user *model.User) error {
	return db.DB.Create(user).Error
}

func GetUserByID(id uint) (*model.User, error) {
	var user model.User
	err := db.DB.First(&user, id).Error
	return &user, err
}

func GetUserList() ([]model.User, error) {
	var users []model.User
	err := db.DB.Find(&users).Error
	return users, err
}

func UpdateUser(user *model.User) error {
	return db.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return db.DB.Delete(&model.User{}, id).Error
}
