package helpers

import "github.com/schweigert/inscreveai/model"

func SignInOrLogin(isAuth bool, userInfo *model.UserInfo) string {
	if isAuth {
		return userInfo.Name
	}

	return "Entrar"
}
