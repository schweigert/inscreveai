package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/controllers/event_controller"
	"github.com/schweigert/inscreveai/controllers/home_controller"
	"github.com/schweigert/inscreveai/controllers/subscription_controller"
	"github.com/schweigert/inscreveai/model"
)

func migrate() {
	db := model.Db()
	defer db.Close()

	db.AutoMigrate(
		&model.UserInfo{},
		&model.Event{},
		&model.Subscription{},
	)
}

func main() {
	migrate()

	r := gin.Default()
	store, _ := sessions.NewRedisStore(10, "tcp", "redis:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.Static("/public", "./public")

	r.GET("/", home_controller.IndexHandler)

	r.GET("/auth", home_controller.AuthHandler)
	r.POST("/event", event_controller.CreateHandler)

	r.POST("/event/:id/delete", event_controller.DeleteHandler)
	r.POST("/event/:id/subscribe", event_controller.SubscribeHandler)

	r.POST("/subscription/:id/accept", subscription_controller.AcceptHandler)
	r.POST("/subscription/:id/reject", subscription_controller.RejectHandler)

	r.Run("0.0.0.0:3000")
}
