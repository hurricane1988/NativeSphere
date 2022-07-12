package controller

import (
	"NativeSphere/service"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

var Pv pv

type pv struct{}

// GetPvs 获取pv列表、支持过滤、排序、分页
func (p *pv) GetPvs(context *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Page       int    `form:"page"`
		Limit      int    `form:"limit"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind参数失败,错误信息， " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Pv.GetPvs(params.FilterName, params.Limit, params.Page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取pv " + params.FilterName + " 成功",
		"data":    data,
	})
}

// GetPvDetail 获取pv详情
func (p *pv) GetPvDetail(context *gin.Context) {
	params := new(struct {
		PvName string `form:"pv_name"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败，错误信息 " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	logger.Info(params)
	data, err := service.Pv.GetPvDetail(params.PvName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取pv " + params.PvName + "详情成功!",
		"data":    data,
	})
}

// DeletePv 删除pv
func (p *pv) DeletePv(context *gin.Context) {
	params := new(struct {
		PvName string `json:"pv_name"`
	})
	// Delete请求，绑定参数方法修改为context.ShouldBindJSON
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败，错误信息， " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	err := service.Pv.DeletePv(params.PvName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "删除pv " + params.PvName + "成功!",
		"data":    nil,
	})
}
