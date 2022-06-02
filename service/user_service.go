package service

import (
	"gin-project/model"
	"gin-project/serializer"
)

type UserService struct {
}

func (service UserService) Login(user model.User) serializer.Response {
	return serializer.Response{
		Code:  200,
		Data:  nil,
		Msg:   "登录成功",
		Error: "",
	}
}

func (service UserService) Register(user model.User) serializer.Response {
	return serializer.Response{
		Code:  200,
		Data:  nil,
		Msg:   "注册成功",
		Error: "",
	}
}
