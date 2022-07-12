package controller

import (
	"NativeSphere/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

var Namespace namespace

type namespace struct{}

// GetNamespaces 获取namespace列表、支持过滤、排序、分页
func (n *namespace) GetNamespaces(context *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Page       int    `form:"page"`
		Limit      int    `form:"limit"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败 " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Namespace.GetNamespaces(params.FilterName, params.Limit, params.Page)
	if err != nil {
		context.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取namespace列表成功!",
		"data":    data,
	})
}

// GetNamespaceDetail 获取namespace详情
func (n *namespace) GetNamespaceDetail(context *gin.Context) {
	params := new(struct {
		NamespaceName string `form:"namespace_name"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error(errors.New("Bind请求参数失败，错误信息 " + err.Error()))
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Namespace.GetNamespaceDetail(params.NamespaceName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取namespace " + params.NamespaceName + "详情成功!",
		"data":    data,
	})
}

// DeleteNamespace 删除namespace
func (n *namespace) DeleteNamespace(context *gin.Context) {
	params := new(struct {
		NamespaceName string `json:"namespace_name"`
	})
	// DELETE请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.Namespace.DeleteNamespace(params.NamespaceName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "删除namespace " + params.NamespaceName + "成功!",
		"data":    nil,
	})
}

// CreateNamespace 创建namespace
func (n *namespace) CreateNamespace(context *gin.Context) {
	var (
		namespaceCreate = new(service.NamespaceCreate)
		err             error
	)
	if err = context.ShouldBindJSON(namespaceCreate); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if err = service.Namespace.CreateNamespace(namespaceCreate); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "创建namespace " + namespaceCreate.Name + "成功",
		"data":    nil,
	})
}
