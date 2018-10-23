package home_controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/login"
	"github.com/schweigert/inscreveai/model"
	"golang.org/x/oauth2"
)

func AuthHandler(c *gin.Context) {
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	if retrievedState != c.Query("state") {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
		return
	}

	tok, err := login.GoogleConfig.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := login.GoogleConfig.Client(oauth2.NoContext, tok)
	userdata, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defer userdata.Body.Close()

	data, err := ioutil.ReadAll(userdata.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	userInfo := model.UserInfoFromJson(data)

	userInfo.FindOrCreate()
	userInfo.Auth(c)
	c.Redirect(http.StatusPermanentRedirect, "/")
}
