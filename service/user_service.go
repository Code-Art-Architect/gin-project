package service

import (
	"github.com/code-art/gin-project/dao"
	"github.com/code-art/gin-project/model"
)

type UserService struct {
}

func (service UserService) Login(user model.User) bool {
	var one model.User = dao.Mgr.SelectOne(user.Username, user.Password)
	if one.Username != "" {
		return true
	}
	return false
}

func (service UserService) Register(user model.User) bool {
	dao.Mgr.AddUser(&user)
	return true
}
