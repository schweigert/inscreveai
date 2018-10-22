package login

import "golang.org/x/oauth2"

var GoogleConfig *oauth2.Config

func init() {
	GoogleConfig = GoogleOAuth2Config()
}
