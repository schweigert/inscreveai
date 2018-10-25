package event_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/model"
)

func DeleteHandler(c *gin.Context) {
	user := model.UserInfoFromSession(c)
	event := &model.Event{}

	if user.IsAuth(c) {
		db := model.Db()
		defer db.Close()

		id := c.Param("id")
		db.First(event, id)

		if event.UserInfoID == user.ID {
			db.Delete(event)
		}
	}

	c.Redirect(http.StatusFound, "/")
}
