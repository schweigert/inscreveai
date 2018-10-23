package login

import (
	"encoding/json"
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Credentials which stores google ids.
type GoogleCredentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

func ReadGoogleCredentials() GoogleCredentials {
	var credentials GoogleCredentials
	file, err := ioutil.ReadFile("./google_creds.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &credentials)

	return credentials
}

func GoogleOAuth2Config() *oauth2.Config {
	credentials := ReadGoogleCredentials()
	config := &oauth2.Config{
		ClientID:     credentials.Cid,
		ClientSecret: credentials.Csecret,
		RedirectURL:  "http://localhost:3000/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return config
}

func GoogleLoginURL(state string) string {
	return GoogleConfig.AuthCodeURL(state)
}
