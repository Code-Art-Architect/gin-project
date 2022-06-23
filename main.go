package main

import (
	_ "github.com/code-art/gin-project/conf"
	"github.com/code-art/gin-project/router"
)

func main() {
	router.Start()
}
