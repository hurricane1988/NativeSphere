package controller

import (
	"NativeSphere/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

var Ingress ingress

type ingress struct{}

// GetIngresses 获取ingress列表、支持过滤、排序、分页
func (i *ingress) GetIngresses(context *gin.Context) {
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
	data, err := service.Ingress.GetIngresses(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取ingress列表清单成功!",
		"data":    data,
	})
}

// GetIngressDetail 获取ingress详情
func (i *ingress) GetIngressDetail(context *gin.Context) {
	params := new(struct {
		IngressName string `form:"ingressName"`
		Namespace   string `form:"namespace"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败,错误信息," + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Ingress.GetIngressDetail(params.IngressName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取ingress " + params.IngressName + "详情成功!",
		"data":    data,
	})
}

// DeleteIngress 删除ingress
func (i *ingress) DeleteIngress(context *gin.Context) {
	params := new(struct {
		IngressName string `form:"ingressName"`
		Namespace   string `form:"namespace"`
	})
	// DELETE请求，绑定参数方法为context.ShouldBindJSON
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error(errors.New("Bind参数失败,错误信息," + err.Error()))
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.Ingress.DeleteIngress(params.IngressName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "删除ingress " + params.IngressName + "成功!",
		"data":    nil,
	})
}

// CreateIngress 创建ingress
func (i *ingress) CreateIngress(context *gin.Context) {
	var (
		ingressCreate = new(service.IngressCreate)
		err           error
	)
	if err = context.ShouldBindJSON(ingressCreate); err != nil {
		logger.Error(errors.New("Bind请求参数失败,错误信息, " + err.Error()))
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if err = service.Ingress.CreateIngress(ingressCreate); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "创建ingress成功!",
		"data":    nil,
	})
}

// UpdateIngress 更新ingress
func (i *ingress) UpdateIngress(context *gin.Context) {
	params := new(struct {
		Namespace string `json:"namespace"`
		Content   string `json:"content"`
	})
	// PUT请求，绑定参数方法改为
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败,错误信息, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.Ingress.UpdateIngress(params.Namespace, params.Content)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "更新ingress成功!",
		"data":    nil,
	})
}
