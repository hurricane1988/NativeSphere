//package middle
//
//import (
//	"github.com/gin-gonic/gin"
//	"k8s-platform/pkg/e"
//	"k8s-platform/utils"
//	"net/http"
//	"time"
//)
//
//// JWTAuth JWTToken中间件，检查token
//func JWTAuth() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		var code int
//		var data interface{}
//
//		code = e.SUCCESS
//		token := context.Query("token")
//		//auth := context.GetHeader("Authorization")
//		//token := strings.Split(auth, "Bearer ")[1]
//		if token == "" {
//			code = e.ERROR_AUTH_CHECK_NO_TOKEN
//		} else {
//			claims, err := utils.ParseToken(token)
//			if err != nil {
//				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
//			} else if time.Now().Unix() > claims.ExpiresAt {
//				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
//			}
//		}
//		if code != e.SUCCESS {
//			context.JSON(http.StatusUnauthorized, gin.H{
//				"code":    code,
//				"message": e.GetMsg(code),
//				"data":    data,
//			})
//			context.Abort()
//			return
//		}
//		context.Next()
//	}
//}

package middle

import (
	"NativeSphere/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWTAuth jwt认证函数
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 对登录接口放行
		if len(context.Request.URL.String()) >= 10 && context.Request.URL.String()[0:10] == "/api/login" {
			context.Next()
		} else {
			// 处理验证逻辑
			token := context.Request.Header.Get("Authorization")
			if token == "" {
				context.JSON(http.StatusBadRequest, gin.H{
					"message": "请求未携带token,无权限访问",
					"data":    nil,
				})
				context.Abort()
				return
			}
			// 解析token内容
			claims, err := utils.JWTToken.ParseToken(token)
			if err != nil {
				// token过期错误
				if err.Error() == "TokenExpired" {
					context.JSON(http.StatusBadRequest, gin.H{
						"message": err.Error(),
						"data":    nil,
					})
					context.Abort()
					return
				}
				// 解析其他错误
				context.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
					"data":    nil,
				})
				context.Abort()
				return
			}
			// 继续交由下一个路由处理,并将解析出的信息传递下去
			context.Set("claims", claims)
			context.Next()
		}
	}
}
