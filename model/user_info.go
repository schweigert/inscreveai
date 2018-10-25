package model

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	gorm.Model
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
	Hd            string `json:"hd"`
	Events        []Event
	Subscriptions []Subscription
}

func (userInfo *UserInfo) Auth(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("auth_token_sub") == nil {
		session.Set("auth_token_sub", userInfo.Sub)
		session.Set("auth_token_mail", userInfo.Email)
		session.Set("auth_token_hash", userInfo.Hash())
		session.Save()
	}
}

func (userInfo *UserInfo) IsAuth(c *gin.Context) bool {
	session := sessions.Default(c)
	if session.Get("auth_token_sub") != nil {
		sub := session.Get("auth_token_sub")
		email := session.Get("auth_token_mail")
		hash := session.Get("auth_token_hash")
		return sub == userInfo.Sub && email == userInfo.Email && hash == userInfo.Hash()
	}
	return false
}

func (userInfo *UserInfo) Hash() string {
	data := userInfo.Email + ":" + userInfo.Sub
	hasher := sha1.New()
	hasher.Write([]byte(data))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (userInfo *UserInfo) FindOrCreate() {
	db := Db()
	db.Where("email = ?", userInfo.Email).First(userInfo)
	if userInfo.ID == 0 {
		db.Create(userInfo)
	}
}

func UserInfoFromSession(c *gin.Context) *UserInfo {
	db := Db()
	defer db.Close()

	session := sessions.Default(c)
	if session.Get("auth_token_sub") != nil {
		sub := session.Get("auth_token_sub")
		email := session.Get("auth_token_mail")
		hash := session.Get("auth_token_hash")

		userInfo := &UserInfo{Email: ""}
		db.Where("sub = ? AND email = ?", sub, email).Preload("Subscriptions").First(userInfo)

		if userInfo.Hash() == hash {
			return userInfo
		}
	}

	return &UserInfo{Email: "sem.email@email.com", EmailVerified: false, Name: "Sem Nome"}
}

func UserInfoFromJson(data []byte) *UserInfo {
	userInfo := &UserInfo{}
	json.Unmarshal(data, userInfo)
	return userInfo
}
