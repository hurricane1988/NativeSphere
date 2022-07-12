package controller

import (
	"NativeSphere/service"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

var DaemonSet daemonSet

type daemonSet struct{}

// GetDaemonSets 获取daemonset列表，支持过滤、排序、分页
func (d *daemonSet) GetDaemonSets(context *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Page       int    `form:"page"`
		Limit      int    `form:"limit"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind参数失败,错误信息 " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	data, err := service.DaemonSet.GetDaemonSets(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取命名空间" + params.Namespace + "下demonSet成功!",
		"data":    data,
	})
}

// GetDaemonSetDetail 获取damonSet详情
func (d *daemonSet) GetDaemonSetDetail(context *gin.Context) {
	params := new(struct {
		DaemonSetName string `form:"daemonset_name"`
		Namespace     string `form:"namespace"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, 错误信息 " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.DaemonSet.GetDaemonSetDetail(params.DaemonSetName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取damonSet " + params.DaemonSetName + "详细成功!",
		"data":    data,
	})
}

// DeleteDaemonSet 删除daemonSet
func (d *daemonSet) DeleteDaemonSet(context *gin.Context) {
	params := new(struct {
		DaemonSetName string `json:"daemonSet_name"`
		Namespace     string `json:"namespace"`
	})
	//DELETE请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	err := service.DaemonSet.DeleteDaemonSet(params.DaemonSetName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "删除DaemonSet成功",
		"data":    nil,
	})
}

// UpdateDaemonSet 更新daemonSet
func (d *daemonSet) UpdateDaemonSet(context *gin.Context) {
	params := new(struct {
		Namespace string `json:"namespace"`
		Content   string `json:"content"`
	})
	//PUT请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	err := service.DaemonSet.UpdateDaemonSet(params.Namespace, params.Content)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "更新DaemonSet成功",
		"data":    nil,
	})
}
