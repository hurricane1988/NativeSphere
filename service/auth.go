package service

import "NativeSphere/config"

// CheckAuth 认证函数方法
func CheckAuth(username, password string) bool {
	if username == config.Username && password == config.Password {
		return true
	} else {
		return false
	}
}
