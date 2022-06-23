package controller

import (
	"github.com/code-art/gin-project/conf"
	"github.com/code-art/gin-project/model"
	"github.com/code-art/gin-project/service"
	"github.com/gin-gonic/gin"
)

var userService service.UserService

func HandleLogin(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	conf.Logrus.Info("用户正在登录...")
	
	var user = model.User{
		Username: username,
		Password: password,
	}

	flag := userService.Login(user)

	if flag {
		conf.Logrus.Info("用户登录成功!")
		context.Redirect(301, "/")
	} else {
		context.JSON(400, gin.H{
			"msg": "登录失败",
		})
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
