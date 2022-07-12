package controller

import (
	"NativeSphere/service"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

var Node node

type node struct{}

// GetNodes 获取node列表，支持过滤、排序、分页
func (n *node) GetNodes(context *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Page       int    `form:"page"`
		Limit      int    `form:"limit"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	data, err := service.Node.GetNodes(params.FilterName, params.Limit, params.Page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"msg":  "获取Node列表成功",
		"data": data,
	})
}

// GetNodeDetail 获取node详情
func (n *node) GetNodeDetail(context *gin.Context) {
	params := new(struct {
		NodeName string `form:"node_name"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	data, err := service.Node.GetNodeDetail(params.NodeName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "获取Node详情成功",
		"data":    data,
	})
}
