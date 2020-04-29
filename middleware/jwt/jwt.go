package jwt

import (
	"blog/pkg/app"
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func Jwt() gin.HandlerFunc  {
	return func(c *gin.Context) {
		appG := app.Gin{c}
		token := c.Request.Header.Get("token")
		if token == "" {
			appG.Response(e.NOT_LOGGED_IN, nil)
			c.Abort()
			return
		}
		code := e.SUCCESS
		claims, err := util.ParseToken(token)
		if err != nil {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}
		if code != e.SUCCESS {
			appG.Response( code, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
