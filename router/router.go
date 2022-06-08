package router

import (
	"gin-project/controller"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**")
	r.Static("/assets", "./assets")

	r.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})
	r.GET("/login.html", func(context *gin.Context) {
		context.HTML(200, "login.html", nil)
	})
	r.GET("/register.html", func(context *gin.Context) {
		context.HTML(200, "register.html", nil)
	})

	r.GET("/user", controller.ListUser)
	r.POST("/register", controller.HandleRegister)
	r.POST("/login", controller.HandleLogin)
	r.Run()
}
