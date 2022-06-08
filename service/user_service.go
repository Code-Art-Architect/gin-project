package service

import (
	"gin-project/dao"
	"gin-project/model"
)

type UserService struct {
}

func (service UserService) Login(user model.User) bool {
	one := dao.Mgr.SelectOne(user.Username, user.Password)
	if one.Username != "" {
		return true
	}
	return false
}

func (service UserService) Register(user model.User) bool {
	dao.Mgr.AddUser(&user)
	return true
}
