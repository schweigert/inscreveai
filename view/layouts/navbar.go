package layouts

import "github.com/schweigert/inscreveai/view/html"

func Navbar(yield func() html.Dom) []byte {
	return Application(
		func() html.Dom {
			return html.DivTag(
				"",
				html.NavTag(
					`class="navbar navbar-expand-lg navbar-dark bg-danger"`,
					html.UlTag(
						`class="navbar-nav mr-auto"`,
						html.LiTag(
							`class="nav-item active"`,
							html.ATag(
								`class="navbar-brand" href="/"`,
								html.ITag(`class="far fa-bookmark fa-lg"`),
								html.Text("Inscreve.ai"),
							),
						),
					),
					html.FormTag(
						`class="form-inline float-right"`,
						html.InputTag(
							`class="form-control mr-sm-2" type="search" placeholder="Busca" aria-label="Busca" name="query"`,
						),
						html.ButtonTag(
							`class="btn btn-outline-white my-sm-0" type="submit"`,
							html.ITag(`class="fas fa-search"`),
						),
						html.ButtonTag(
							`class="btn btn-outline-warning ml-2 my-sm-0"`,
							html.ITag(`class="fab fa-google"`),
						),
					),
				),
				yield(),
			)
		},
	)
}
