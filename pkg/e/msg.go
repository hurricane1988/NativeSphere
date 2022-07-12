package e

// MsgFlags 定义错误代码注释文本
var MsgFlags = map[int]string{
	SUCCESS:                        "操作成功",
	ERROR:                          "操作失败",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_AUTH_CHECK_NO_TOKEN:      "请求未携带Token",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "生成Token的请求账号、密码错误",
}

// GetMsg 自定义返回信息函数方法
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
