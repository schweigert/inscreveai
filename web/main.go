package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/controllers/home_controller"
)

func main() {
	r := gin.Default()

	store, err := sessions.NewRedisStore(10, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}

	r.Use(sessions.Sessions("session", store))

	r.Static("/public", "./public")

	r.GET("/", home_controller.IndexHandler)
	r.GET("/auth", home_controller.AuthHandler)

	r.Run("0.0.0.0:3000")
}
