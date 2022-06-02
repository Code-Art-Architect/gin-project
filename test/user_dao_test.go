package test

import (
	"gin-project/dao"
	"gin-project/model"
	"testing"
)

func TestAddUser(t *testing.T) {
	user := model.User{
		Username: "洛必达",
		Password: "123456",
	}

	dao.Mgr.AddUser(&user)
}
