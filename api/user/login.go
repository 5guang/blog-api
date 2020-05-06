package user

import (
	"blog/models/request"
	"blog/models/response"
	"blog/pkg/app"
	"blog/pkg/e"
	"blog/service/auth_service"
	"blog/service/user_service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context)  {
	var (
		appG = app.Gin{C:c}
		reqLogin request.ReqLogin
		resLoginData response.ResLoginData
	)
	errCode := app.BindAndValid(c, &reqLogin)
	if errCode != e.SUCCESS {
		appG.Response( errCode, nil)
		return
	}

	userService := user_service.User{
		Username: reqLogin.Body.Username,
		Password: reqLogin.Body.Password,
		AdminPassword: reqLogin.Body.AdminPassword,
	}
	errCode, resLoginData = userService.Login()

	token, code := auth_service.GetAuthInfo(reqLogin.Body.Username, reqLogin.Body.Password)
	if code != e.SUCCESS {
		appG.Response( code, nil)
		return
	}
	c.Request.Header.Set("token", token)

	appG.Response( errCode, resLoginData)
}
