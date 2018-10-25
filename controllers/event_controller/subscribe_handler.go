package event_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/model"
)

func SubscribeHandler(c *gin.Context) {
	user := model.UserInfoFromSession(c)

	if user.IsAuth(c) {
		db := model.Db()
		defer db.Close()

		event := model.Event{}

		id := c.Param("id")
		db.Find(&event, id)

		if (event.ID != model.Event{}.ID) {
			subscription := &model.Subscription{UserInfoID: user.ID, EventID: event.ID, Waiting: true, Confirmed: false}
			db.Create(subscription)
		}
	}

	c.Redirect(http.StatusFound, "/")
}
