package auth_service

import (
	"blog/dao"
	"blog/pkg/e"
	"blog/pkg/util"
)

func GetAuthInfo(username, password string) (string, int) {
	isExist := dao.CheckAuth(username)
	token := ""
	code := e.SUCCESS
	var err error
	if isExist == true {
		token, err = util.GenerateToken(username, password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			code = e.SUCCESS
		}
	} else {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
	}
	return token, code
}
