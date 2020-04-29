package api

import (
	"blog/models/request"
	"blog/pkg/app"
	"blog/pkg/e"
	"blog/service/auth_service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context)  {
	var (
		appG = app.Gin{C:c}
		reqLogin request.ReqLogin
		)
	 errCode := app.BindAndValid(c, &reqLogin)
	if errCode != e.SUCCESS {
		appG.Response( errCode, nil)
		return
	}
	token, code := auth_service.GetAuthInfo(reqLogin)
	c.Request.Header.Set("token", token)

	appG.Response( code, nil)
}
