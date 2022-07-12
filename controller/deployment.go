package controller

import (
	"NativeSphere/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"strconv"
)

var Deployment deployment

type deployment struct{}

// GetDeployments 获取deployment列表，支持过滤，排序、分页
func (d *deployment) GetDeployments(context *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Page       int    `form:"page"`
		Limit      int    `form:"limit"`
	})

	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Deployment.GetDeployments(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取namespace" + params.Namespace + "Deployment列表成功",
		"data":    data,
	})
}

// GetDeploymentDetail 获取deployment详情
func (d *deployment) GetDeploymentDetail(context *gin.Context) {
	params := new(struct {
		DeploymentName string `form:"deployment_name"`
		Namespace      string `form:"namespace"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败,错误信息 " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Deployment.GetDeploymentDetail(params.DeploymentName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取namespace " + params.Namespace + "成功",
		"data":    data,
	})
}

// CreateDeployment 创建deployment
func (d *deployment) CreateDeployment(context *gin.Context) {
	var (
		deployCreate = new(service.DeployCreate)
		err          error
	)
	if err = context.ShouldBindJSON(deployCreate); err != nil {
		logger.Error("Bind请求参数失败，" + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if err = service.Deployment.CreateDeployment(deployCreate); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "创建deployment" + deployCreate.Name + "成功",
		"data":    nil,
	})
}

// ScaleDeployment 设置deployment副本数
func (d *deployment) ScaleDeployment(context *gin.Context) {
	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
		ScaleNum       int    `json:"scale_num"`
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
	data, err := service.Deployment.ScaleDeployment(params.DeploymentName, params.Namespace, params.ScaleNum)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "设置deployment" + params.DeploymentName + "副本数至" + strconv.Itoa(params.ScaleNum) + "成功",
		"data":    fmt.Sprintf("最新副本数：%d\n", data),
	})
}

// DeleteDeployment 删除deployment
func (d *deployment) DeleteDeployment(context *gin.Context) {
	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
	})
	// DELETE请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := context.ShouldBindJSON(params); err != nil {
		logger.Error("Bind请求参数失败, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.Deployment.DeleteDeployment(params.DeploymentName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "删除deployment" + params.DeploymentName + "成功!",
		"data":    nil,
	})
}

// RestartDeployment 重启deployment
func (d *deployment) RestartDeployment(context *gin.Context) {
	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
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
	err := service.Deployment.RestartDeployment(params.DeploymentName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "重启Deployment成功",
		"data":    nil,
	})
}

// UpdateDeployment 更新deployment
func (d *deployment) UpdateDeployment(context *gin.Context) {
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
	err := service.Deployment.UpdateDeployment(params.Namespace, params.Content)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "更新Deployment成功",
		"data":    nil,
	})
}

// GetDeployReplicaSets 获取deployment的历史版本信息
func (d *deployment) GetDeployReplicaSets(context *gin.Context) {
	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
	})
	if err := context.Bind(params); err != nil {
		logger.Error("Bind请求参数失败,错误信息, " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.Deployment.GetDeployReplicaSets(params.DeploymentName, params.Namespace)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取deployment " + params.DeploymentName + "历史版本信息成功!",
		"data":    nil,
	})
}

// GetDeployNumPerNp 获取每个namespace的pod数量
func (d *deployment) GetDeployNumPerNp(context *gin.Context) {
	data, err := service.Deployment.GetDeployNumPerNp()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "获取每个namespace的deployment数量成功",
		"data":    data,
	})
}
