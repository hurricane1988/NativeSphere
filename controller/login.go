package controller

import (
	"NativeSphere/service"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

var Login login

type login struct{}

// Auth 验证账号密码
func (login *login) Auth(context *gin.Context) {
	params := new(struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	})
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	err := service.Login.Auth(params.UserName, params.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"data":    nil,
	})
}
