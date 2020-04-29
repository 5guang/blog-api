package user

import (
	"blog/models/request"
	"blog/pkg/app"
	"blog/pkg/e"
	"blog/service/user_service"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context)  {
	appG := app.Gin{C: c}

	var (
		reqRegister request.ReqRegister
	)

	errCode := app.BindAndValid(c,&reqRegister)
	if errCode != e.SUCCESS {
		appG.Response(errCode, nil)
		return
	}
	userService := user_service.User{
		Username: reqRegister.Body.Username,
		Password: reqRegister.Body.Password,
		NickName: reqRegister.Body.Nickname,
		Email: reqRegister.Body.Email,
	}
	errCode = userService.Register()
	appG.Response(errCode, nil)

}