package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var TestApi testApi

type testApi struct{}

func (t *testApi) TestAPI(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg":  "test success",
		"data": nil,
	})
}
