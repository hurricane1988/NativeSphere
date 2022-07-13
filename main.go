package main

import (
	"NativeSphere/config"
	"NativeSphere/controller"
	"NativeSphere/db"
	"NativeSphere/middle"
	"NativeSphere/service"
	"NativeSphere/utils"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

// 设置打印格式信息
var (
	Yellow       = color.New(color.FgHiYellow, color.Bold).SprintFunc()
	YellowItalic = color.New(color.FgHiYellow, color.Bold, color.Italic).SprintFunc()
	Green        = color.New(color.FgHiGreen, color.Bold).SprintFunc()
	Blue         = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	Cyan         = color.New(color.FgCyan, color.Bold, color.Underline).SprintFunc()
	Red          = color.New(color.FgHiRed, color.Bold).SprintFunc()
	White        = color.New(color.FgWhite).SprintFunc()
	WhiteBold    = color.New(color.FgWhite, color.Bold).SprintFunc()
	forceDetail  = "yaml"
)

func main() {
	// 初始化k8s client
	service.K8s.Init() // 可以使用service.K8s.ClientSet挎包调用
	// 初始化数据库
	db.Init()
	// 初始化gin对象
	router := gin.Default()
	// 获取token路由
	router.GET("/auth", controller.GetAuth)
	// 加载jwt中间件
	router.Use(middle.JWTAuth())
	// router := gin.New()
	// 跨域配置(中间需要在初始化路由之前配置)
	router.Use(middle.Cores())
	// 设置gin运行模式
	gin.SetMode(config.GinMode)
	// 挎包调用router的初始化方法
	controller.Router.InitApiRouter(router)
	// 打印彩色终端
	utils.PrintColor()
	// gin程序启动
	err := router.Run(config.ListenAddr)
	if err != nil {
		return
	}
	// 关闭db数据库连接
	db.Close()
}
