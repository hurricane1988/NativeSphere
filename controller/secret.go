package controller

import (
	"NativeSphere/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

var Secret secret

type secret struct{}

// GetSecrets 获取secret列表、支持过滤、排序、分页
func (s *secret) GetSecrets(context *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Page       int    `form:"page"`
		Limit      int    `form:"limit"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind参数失败,错误信息," + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Secret.GetSecrets(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取secret " + params.FilterName + "成功!",
		"data":    data,
	})
}

// GetSecretDetail 获取secret详情
func (s *secret) GetSecretDetail(context *gin.Context) {
	params := new(struct {
		SecretName string `form:"secretName"`
		Namespace  string `form:"namespace"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error(errors.New("获取secret " + params.SecretName + "详细失败,错误信息, " + err.Error()))
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Secret.GetSecretDetail(params.SecretName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    data,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取Secret " + params.SecretName + "成功!",
		"data":    data,
	})
}

// DeleteSecret 删除secret
func (s *secret) DeleteSecret(context *gin.Context) {
	params := new(struct {
		SecretName string `json:"secretName"`
		Namespace  string `json:"namespace"`
	})
	// Delete请求，绑定参数方法为context.ShouldBindJSON
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error("Bind参数失败,错误信息, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.Secret.DeleteSecret(params.SecretName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "删除secret " + params.SecretName + "成功!",
		"data":    nil,
	})
}

// UpdateSecret 更新secret
func (s *secret) UpdateSecret(context *gin.Context) {
	params := new(struct {
		Namespace string `json:"namespace"`
		Content   string `json:"content"`
	})
	// PUT请求,绑定参数方法为context.ShouldBindJSON
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error("Bind参数失败,错误信息, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.Secret.UpdateSecret(params.Namespace, params.Content)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "更新secret成功!",
		"data":    nil,
	})
}
