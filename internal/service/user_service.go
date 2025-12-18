package service

import (
	"crud/internal/dao"
	"crud/internal/model"
	"crud/pkg/cache"
	"crud/pkg/password"
	"encoding/json"
	"errors"
	"time"
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
	//1.查redis
	key := cache.UserKey(id)
	val, err := cache.RDB.Get(cache.Ctx, key).Result()
	if err == nil {
		var user model.User
		_ = json.Unmarshal([]byte(val), &user)
		return &user, nil
	}
	//2.查数据库
	user, err := dao.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	//3.写入redis
	bytes, _ := json.Marshal(user)
	cache.RDB.Set(cache.Ctx, key, bytes, 5*time.Minute)

	return user, nil
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
	if err := dao.UpdateUser(user); err != nil {
		return err
	}

	//删除缓存
	cache.RDB.Del(cache.Ctx, cache.UserKey(id))

	return nil
}

func DeleteUser(id uint) error {
	if err := dao.DeleteUser(id); err != nil {
		return err
	}

	cache.RDB.Del(cache.Ctx, cache.UserKey(id))
	return nil
}
