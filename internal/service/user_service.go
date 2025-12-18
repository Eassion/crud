package service

import (
	"crud/internal/dao"
	"crud/internal/model"
	"crud/pkg/password"
	"errors"
)

func CreateUser(username, passwordStr string, age int) error {
	if username == "" || passwordStr == "" {
		return errors.New("用户名或密码不能为空")
	}

	hashPwd, err := password.HashPassword(passwordStr)
	if err != nil {
		return errors.New("密码加密失败")
	}

	user := &model.User{
		Username: username,
		Password: hashPwd,
		Age:      age,
	}

	return dao.CreateUser(user)
}

func GetUser(id uint) (*model.User, error) {
	return dao.GetUserByID(id)
}

func GetUserList() ([]model.User, error) {
	return dao.GetUserList()
}

func UpdateUser(id uint, age int) error {
	user, err := dao.GetUserByID(id)
	if err != nil {
		return err
	}

	user.Age = age
	return dao.UpdateUser(user)
}

func DeleteUser(id uint) error {
	return dao.DeleteUser(id)
}
