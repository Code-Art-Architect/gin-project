package main

import (
	"gin-project/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v := r.Group("/api/v1")
	{
		v.POST("/user/login", controller.HandleLogin)
		v.POST("/user/register", controller.HandleRegister)
	}

	_ = r.Run()
}
