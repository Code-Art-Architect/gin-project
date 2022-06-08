package controller

import (
	"gin-project/model"
	"gin-project/service"
	"github.com/gin-gonic/gin"
)

var userService service.UserService

func HandleLogin(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	var user = model.User{
		Username: username,
		Password: password,
	}

	flag := userService.Login(user)

	if flag {
		context.Redirect(301, "/")
	} else {
		context.HTML(400, "/login.html", "登录失败!")
	}
}

func HandleRegister(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	var user = model.User{
		Username: username,
		Password: password,
	}

	flag := userService.Register(user)

	if flag {
		context.Redirect(301, "/login.html")
	}
}

func ListUser(context *gin.Context) {
	context.HTML(200, "user.html", nil)

}
