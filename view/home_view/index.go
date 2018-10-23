package home_view

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/model"
	"github.com/schweigert/inscreveai/view/html"
	"github.com/schweigert/inscreveai/view/layouts"
)

func Index(c *gin.Context) []byte {
	user := model.UserInfoFromSession(c)
	name := "Sem nome"
	if user != nil {
		name = user.Name
	}
	return layouts.Navbar(
		c,
		func() html.Dom {
			return html.DivTag(
				"",
				html.Text(name),
			)
		},
	)
}
