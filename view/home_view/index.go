package home_view

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/model"
	"github.com/schweigert/inscreveai/view/html"
	"github.com/schweigert/inscreveai/view/layouts"
)

func Index(isAuth bool, currentUser *model.UserInfo, events []html.Dom, c *gin.Context) []byte {
	return layouts.Navbar(
		isAuth, currentUser, c,
		func() html.Dom {
			return html.DivTag(
				`class="container mt-5"`,
				html.DivTag(
					`clas="row"`,
					html.DivTag(
						`class="col-sm-12"`,
						html.DivTag(
							`class="card"`,
							html.DivTag(
								`class="card-header text-center"`,
								html.Text("Crie o evento dos teus sonhos!"),
							),
							html.DivTag(
								`class="card-body"`,
								html.FormTag(
									`action="/event" method="POST"`,
									html.DivTag(
										`class="row"`,
										html.LabelTag(
											`for="name"`,
											html.Text("Nome:"),
										),
										html.InputTag(
											`id="name" name="name" type="text" class="form-control" placeholder="Um nome fabuloso!"`,
										),
									),
									html.DivTag(
										`class="row"`,
										html.LabelTag(
											`for="description"`,
											html.Text("Descrição:"),
										),
										html.InputTag(
											`id="description" name="description" type="text" class="form-control" placeholder="Uma descrição atrativa!"`,
										),
									),
									html.DivTag(
										`class="text-center mt-3 mb-3"`,
										html.ButtonTag(
											`class="btn btn-outline-danger" type="submit"`,
											html.ITag(`class="fas fa-calendar"`),
											html.Text("Criar Evento!"),
										),
									),
								),
							),
						),
					),
				),
				html.DivTag(
					`class="row"`,
					events...,
				),
			)
		},
	)
}
