package controller

import (
	"NativeSphere/service"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

var ConfigMap configMap

type configMap struct{}

// GetConfigMaps 获取configMap列表、支持过滤、排序、分页
func (c *configMap) GetConfigMaps(context *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Page       int    `form:"page"`
		Limit      int    `form:"limit"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败，错误信息, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	data, err := service.ConfigMap.GetConfigMaps(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取ConfigMap列表成功!",
		"data":    data,
	})
}

// GetConfigMapDetail 获取configMap详情
func (c *configMap) GetConfigMapDetail(context *gin.Context) {
	params := new(struct {
		ConfigMapName string `form:"configmap_name"`
		Namespace     string `form:"namespace"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.ConfigMap.GetConfigMapDetail(params.ConfigMapName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取ConfigMap " + params.ConfigMapName + "详情成功!",
		"data":    data,
	})
}

// UpdateConfigMap 更新ConfigMap
func (c *configMap) UpdateConfigMap(context *gin.Context) {
	params := new(struct {
		Namespace string `json:"namespace"`
		Content   string `json:"content"`
	})
	// PUT请求，绑定参数方法为context.ShouldBindJSON
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败,错误信息 " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.ConfigMap.UpdateConfigMap(params.Namespace, params.Content)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "更新configMap成功!",
		"data":    nil,
	})
}

// DeleteConfigMap 删除configMap
func (c *configMap) DeleteConfigMap(context *gin.Context) {
	params := new(struct {
		ConfigMapName string `json:"configMapName"`
		Namespace     string `json:"namespace"`
	})
	// DELETE请求，绑定参数方法为context.ShouldBindJSON
	if err := context.ShouldBindJSON(params); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.ConfigMap.DeleteConfigMap(params.ConfigMapName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "删除configMap " + params.ConfigMapName + "成功!",
		"data":    nil,
	})
}
