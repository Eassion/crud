package service

import (
	"crud/internal/dao"
	"crud/pkg/jwt"
	"crud/pkg/password"
	"errors"
)

func Login(username, passwordStr string) (string, error) {
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	if !password.CheckPassword(user.Password, passwordStr) {
		return "", errors.New("用户名或密码错误")
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("token generate failed")
	}

	return token, nil
} // 工程点：不告诉前端是用户名错还是密码错
