package subscription_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/model"
)

func AcceptHandler(c *gin.Context) {
	user := model.UserInfoFromSession(c)
	sub := model.Subscription{}
	subId := c.Param("id")
	event := model.Event{}

	if user.IsAuth(c) {
		db := model.Db()
		defer db.Close()
		db.Find(&sub, subId)
		db.Model(&sub).Related(&event)

		if event.UserInfoID == user.ID {
			sub.Waiting = false
			sub.Confirmed = true
			db.Save(&sub)
		}
	}

	c.Redirect(http.StatusFound, "/")
}
