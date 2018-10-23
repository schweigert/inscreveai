package home_view

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/model"
	"github.com/schweigert/inscreveai/view/html"
	"github.com/schweigert/inscreveai/view/layouts"
)

func Index(isAuth bool, currentUser *model.UserInfo, c *gin.Context) []byte {
	return layouts.Navbar(
		isAuth, currentUser, c,
		func() html.Dom {
			return html.DivTag(
				"",
				html.Text(currentUser.Name),
			)
		},
	)
}
