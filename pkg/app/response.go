package app

import (
	"blog/models/response"
	"blog/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response( errCode int, data interface{})  {
	g.C.JSON(http.StatusOK, response.CommonResponse{
		Head: response.Head{Token: g.C.Request.Header.Get("token")},
		Body: response.Body{
			Status: response.Status{
				Code:errCode,
				Msg: e.GetMsg(errCode),
			},
			Data: data,
		},
	})
}
