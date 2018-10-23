package event_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/model"
)

func CreateHandler(c *gin.Context) {
	description := c.PostForm("description")
	name := c.PostForm("name")

	user := model.UserInfoFromSession(c)

	if user.IsAuth(c) {
		db := model.Db()
		defer db.Close()

		db.Create(&model.Event{Description: description, Name: name, UserInfoId: user.ID})
	}

	c.Redirect(http.StatusFound, "/")
}
