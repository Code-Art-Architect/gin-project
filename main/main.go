package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})

	_ = r.Run()
}
