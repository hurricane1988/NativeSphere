package controller

import (
	"NativeSphere/config"
	"NativeSphere/pkg/e"
	"NativeSphere/service"
	"NativeSphere/utils"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetAuth 获取认证信息
func GetAuth(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")

	valid := validation.Validation{}

	a := config.Auth{
		Username: username,
		Password: password,
	}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		isExist := service.CheckAuth(username, password)
		if isExist {
			token, err := utils.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data": gin.H{
			"token": data,
		},
	})
}
