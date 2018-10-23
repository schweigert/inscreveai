package helpers

import (
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/login"
	"github.com/schweigert/inscreveai/model"
	"github.com/schweigert/inscreveai/view/html"
)

func GoogleOAuthState(c *gin.Context) string {
	state := login.RandToken()
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()

	return state
}

func AuthButton(isAuth bool, currentUser *model.UserInfo, c *gin.Context) html.Dom {
	return html.If(
		isAuth,
		html.ATag(
			fmt.Sprintf(`class="btn btn-outline-warning ml-2 my-sm-0" href="%s"`, login.GoogleLoginURL(GoogleOAuthState(c))),
			html.ITag(`class="fab fa-google"`),
			html.Text(currentUser.Name),
		),
		html.ATag(
			fmt.Sprintf(`class="btn btn-outline-warning ml-2 my-sm-0" href="%s"`, login.GoogleLoginURL(GoogleOAuthState(c))),
			html.ITag(`class="fab fa-google"`),
			html.Text("Entrar"),
		),
	)
}
