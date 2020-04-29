package app

import (
	"blog/pkg/e"
	"blog/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValid(c *gin.Context, json interface{}) ( int) {
	err := c.ShouldBind(json)
	if err != nil {
		return  e.INVALID_PARAMS
	}
	validate := validator.New()
	errMsg := validate.Struct(json)
	if errMsg != nil {
		logging.Info(errMsg)
		return e.INVALID_PARAMS
	}

	return  e.SUCCESS
}
