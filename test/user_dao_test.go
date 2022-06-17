package test

import (
	"testing"

	"github.com/code-art/gin-project/dao"
	"github.com/code-art/gin-project/model"
)

func TestAddUser(t *testing.T) {
	user := model.User{
		Username: "洛必达",
		Password: "123456",
	}

	dao.Mgr.AddUser(&user)
}
