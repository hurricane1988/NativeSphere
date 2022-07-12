package service

import (
	"NativeSphere/config"
	"errors"
	"github.com/wonderivan/logger"
)

var Login login

type login struct{}

// Auth 验证账号密码
func (login *login) Auth(username, password string) (err error) {
	if username == config.AdminUser && password == config.AdminPassword {
		return nil
	} else {
		logger.Error("登录失败,用户名或密码错误")
		return errors.New("登录失败,用户名或密码错误")
	}
	return nil
}
