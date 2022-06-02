package controller

import (
	"gin-project/dao"
	"gin-project/model"
	"gin-project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userService service.UserService

func HandleLogin(context *gin.Context) {
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数有误!",
		})
	} else {
		response := userService.Login(user)
		context.JSON(200, response)
	}
}

func HandleRegister(context *gin.Context) {
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数有误!",
		})
	} else {
		response := userService.Register(user)
		context.JSON(200, response)
	}
}

func AddUser(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(&user)
}
