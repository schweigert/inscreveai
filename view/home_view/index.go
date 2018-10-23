package home_view

import (
	"github.com/schweigert/inscreveai/view/html"
	"github.com/schweigert/inscreveai/view/layouts"
)

func Index() []byte {
	return layouts.Navbar(
		func() html.Dom {
			return html.DivTag(
				"",
				html.Text("inscreve.ai"),
			)
		},
	)
}
