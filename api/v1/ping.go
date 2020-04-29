package v1

import (
	"blog/pkg/app"
	"blog/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Summary 检查
// @Produce  json
// @Success 200 {object} response.CommonResponse
// @Router /ping [get]
func Ping(c *gin.Context) {
	app := app.Gin{c}
	app.Response( e.PING, nil)
}
