package controller

import (
	"NativeSphere/service"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

var Servicev1 servicev1

type servicev1 struct{}

// GetServices 获取service列表、支持过滤、分页
func (s *servicev1) GetServices(context *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Page       int    `form:"page"`
		Limit      int    `form:"limit"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败,错误信息," + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Servicev1.GetServices(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取Service列表成功!",
		"data":    data,
	})
}

// GetServiceDetail 获取service详情
func (s *servicev1) GetServiceDetail(context *gin.Context) {
	params := new(struct {
		ServiceName string `form:"service_name"`
		Namespace   string `form:"namespace"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败,错误信息, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Servicev1.GetServiceDetail(params.ServiceName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取Service详情成功!",
		"data":    data,
	})
}

// CreateService 创建service
func (s *servicev1) CreateService(context *gin.Context) {
	var (
		serviceCreate = new(service.ServiceCreate)
		err           error
	)
	if err = context.ShouldBindJSON(serviceCreate); err != nil {
		logger.Error("Bind请求参数失败,错误信息," + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if err = service.Servicev1.CreateService(serviceCreate); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "创建Service成功!",
		"data":    nil,
	})
}

// DeleteService 删除service
func (s *servicev1) DeleteService(context *gin.Context) {
	params := new(struct {
		ServiceName string `json:"serviceName"`
		Namespace   string `json:"namespace"`
	})
	// DELETE请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败,错误信息, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "删除Service成功!",
		"data":    nil,
	})
}

// UpdateService 更新service
func (s *servicev1) UpdateService(context *gin.Context) {
	params := new(struct {
		Namespace string `json:"namespace"`
		Content   string `json:"content"`
	})
	// PUT请求，绑定参数方法为context.ShouldBindJSON
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败,错误信息," + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.Servicev1.UpdateService(params.Namespace, params.Content)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "更新Service成功!",
		"data":    nil,
	})
}
