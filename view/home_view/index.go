package home_view

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/view/html"
	"github.com/schweigert/inscreveai/view/layouts"
)

func Index(c *gin.Context) []byte {
	return layouts.Navbar(
		c,
		func() html.Dom {
			return html.DivTag(
				"",
				html.Text("inscreve.ai"),
			)
		},
	)
}
