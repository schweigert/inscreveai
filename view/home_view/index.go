package home_view

import (
	"github.com/schweigert/inscreveai/view/html"
	"github.com/schweigert/inscreveai/view/layouts"
)

func Index() string {
	return layouts.Application(

		func() html.Dom {
			return html.DivTag(
				"",
				html.Text("inscreve.ai"),
			)
		},
	)
}
