package config

import (
	"github.com/fatih/color"
	"time"
)

const (
	ListenAddr = "0.0.0.0:8080"
	Kubeconfig = "/Users/hurricane/.kube/config"

	// Kubeconfig = "/Users/admin/.kube/config"
	// PodLogTailLine tail 的日志行数

	PodLogTailLine = 2000
)

// JWTSecret 认证secret
const (
	JWTSecret  = "kubeSphere"
	ExpireTime = 10
	Issuer     = "hurricane"
	Username   = "admin"
	Password   = "Harbor12345"
)

// 定义错误代码常量
const (
	SUCCESS                   = 200
	ERROR                     = 500
	INVALID_PARAMS            = 400
	ERROR_AUTH_CHECK_NO_TOKEN = 401

	ERROR_EXIST_TAG         = 10001
	ERROR_NOT_EXIST_TAG     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

// GinMode 设置gin配置
const (
	GinMode = "debug" // 用于测试环境

	// GinMode = "release" // 用于生成环境
	ReadTimeout  = 60
	WriteTimeout = 60
)

// 数据库配置
const (
	DbType     = "mysql"
	DbHost     = "192.168.3.231"
	DbPort     = 3306
	DbName     = "k8s"
	DbUser     = "root"
	DbPassword = "root"
	LogMode    = true
	/* 连接池配置 */
	MaxIdleConns = 10               //最大空闲连接
	MaxOpenConns = 100              //最大连接数
	MaxLifeTime  = 30 * time.Second //最大生存时间
)

// 登录账号
const (
	AdminUser     = "admin"
	AdminPassword = "kubesphere"
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

// websocket全局配置
const (
	END_OF_TRANSMISSION = "\u0004"
)
