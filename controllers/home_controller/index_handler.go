package home_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/model"
	"github.com/schweigert/inscreveai/view/home_view"
)

func IndexHandler(c *gin.Context) {
	user := model.UserInfoFromSession(c)
	isAuth := user.IsAuth(c)
	events := model.AllEventsCards()
	c.Data(http.StatusOK, "text/html; charset=utf-8", home_view.Index(isAuth, user, events, c))
}
