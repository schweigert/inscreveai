package main

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/controllers/home_controller"
)

func main() {
	r := gin.Default()
	r.Static("/public", "./public")

	r.GET("/", home_controller.IndexHandler)

	r.Run("0.0.0.0:3000")
}
