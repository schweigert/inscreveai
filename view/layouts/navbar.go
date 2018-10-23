package layouts

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/login"
	"github.com/schweigert/inscreveai/model"
	"github.com/schweigert/inscreveai/view/html"
)

func googleOAuthState(c *gin.Context) string {
	state := login.RandToken()
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()

	return state
}

func Navbar(isAuth bool, currentUser *model.UserInfo, c *gin.Context, yield func() html.Dom) []byte {
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
						html.If(
							isAuth,
							html.ATag(
								`class="btn btn-outline-warning ml-2 my-sm-0" href="`+login.GoogleLoginURL(googleOAuthState(c))+`"`,
								html.ITag(`class="fab fa-google"`),
								html.Text(currentUser.Name),
							),
							html.ATag(
								`class="btn btn-outline-warning ml-2 my-sm-0" href="`+login.GoogleLoginURL(googleOAuthState(c))+`"`,
								html.ITag(`class="fab fa-google"`),
								html.Text("Entrar"),
							),
						),
					),
				),
				yield(),
			)
		},
	)
}
