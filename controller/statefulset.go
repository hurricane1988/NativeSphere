package controller

import (
	"NativeSphere/service"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

var StatefulSet statefulSet

type statefulSet struct{}

// GetStatefulSets 获取statefulSet列表，支持过滤、排序、分页
func (s *statefulSet) GetStatefulSets(context *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Page       int    `form:"page"`
		Limit      int    `form:"limit"`
	})
	if err := context.Bind(params); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.StatefulSet.GetStatefulSets(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取statefulSet列表成功",
		"data":    data,
	})
}

// GetStatefulSetDetail 获取statefulSet详情
func (s *statefulSet) GetStatefulSetDetail(context *gin.Context) {
	params := new(struct {
		StatefulSetName string `form:"statefulset_name"`
		Namespace       string `form:"namespace"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind参数失败，错误信息" + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.StatefulSet.GetStatefulSetDetail(params.StatefulSetName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "获取StatefulSet详情成功",
		"data":    data,
	})
}

// DeleteStatefulSet 删除statefulset
func (s *statefulSet) DeleteStatefulSet(ctx *gin.Context) {
	params := new(struct {
		StatefulSetName string `json:"statefulset_name"`
		Namespace       string `json:"namespace"`
	})
	//DELETE请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	err := service.StatefulSet.DeleteStatefulSet(params.StatefulSetName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "删除StatefulSet成功",
		"data":    nil,
	})
}

// UpdateStatefulSet 更新statefulSet
func (s *statefulSet) UpdateStatefulSet(context *gin.Context) {
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

	err := service.StatefulSet.UpdateStatefulSet(params.Namespace, params.Content)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "更新StatefulSet成功",
		"data":    nil,
	})
}
