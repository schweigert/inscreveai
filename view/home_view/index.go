package home_view

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/model"
	"github.com/schweigert/inscreveai/view/html"
	"github.com/schweigert/inscreveai/view/layouts"
)

func Index(isAuth bool, currentUser *model.UserInfo, events []html.Dom, adminEvents []html.Dom, c *gin.Context) []byte {
	return layouts.Navbar(
		isAuth, currentUser, c,
		func() html.Dom {
			return html.DivTag(
				`class="container p-5" style="background-color: RGBA(255,255,255,0.8)"`,
				html.If(
					isAuth,
					html.Append(
						html.DivTag(
							`class="row"`,
							html.HTag(
								`1`,
								``,
								html.Text("Crie o seu evento!"),
							),
						),
						html.DivTag(
							`clas="row"`,
							html.DivTag(
								`class="col-sm-12"`,
								html.DivTag(
									`class="card"`,
									html.DivTag(
										`class="card-header text-center"`,
										html.StrongTag(
											``,
											html.SafeText(currentUser.GivenName),
											html.Text(", você pode criar o seu evento!"),
										),
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
												html.TextAreaTag(
													`id="description" name="description" type="textarea" class="form-control" placeholder="Uma descrição atrativa!"`,
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
					),
				),
				html.If(
					isAuth,
					html.Append(
						html.DivTag(
							`class="row"`,
							html.HTag(
								`1`,
								``,
								html.Text("Seus eventos"),
							),
						),
						html.DivTag(
							`class="row"`,
							adminEvents...,
						),
					),
				),
				html.If(
					isAuth,
					html.Append(
						html.DivTag(
							`class="row"`,
							html.HTag(
								`1`,
								``,
								html.Text("Suas inscrições"),
							),
						),
						html.DivTag(
							`class="row"`,
							model.PenddingSubscriptionCards(currentUser.Subscriptions)...,
						),
					),
				),
				html.Append(
					html.DivTag(
						`class="row"`,
						html.HTag(
							`1`,
							``,
							html.Text("Eventos disponíveis"),
						),
					),
					html.DivTag(
						`class="row"`,
						events...,
					),
				),
			)
		},
	)
}
