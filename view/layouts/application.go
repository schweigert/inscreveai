package layouts

import "github.com/schweigert/inscreveai/view/html"

func Application(yield func() html.Dom) string {
	return html.HtmlTag(
		"",
		html.HeadTag(
			"",
			html.TitleTag(
				"",
				html.Text("InscreveAi"),
			),
			html.Text(`<link rel="stylesheet" href="./public/css/bootstrap.css"`),
		),
		html.BodyTag(
			"",
			yield(),
		),
		html.FooterTag(
			"",
			html.ScriptTag(`src="./public/js/jquery.js"`),
			html.ScriptTag(`src="./public/js/popper.js"`),
			html.ScriptTag(`src="./public/js/bootstrap.js"`),
		),
	).HtmlSafe()
}
