package db

import (
	"NativeSphere/config"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"                  //gorm库
	_ "github.com/jinzhu/gorm/dialects/mysql" //gorm对应的mysql驱动
	"github.com/wonderivan/logger"
	"strconv"
	"time"
)

// 初始化数据库变量
var (
	isInit bool
	GORM   *gorm.DB
	err    error
)

// Init db的初始化函数，与数据库建立连接
func Init() {
	// 判断是否已经初始化了
	if isInit {
		return
	}
	// 组装连接配置
	// parseTime是查询结果是否自动解析为时间
	// loc是MySQL的时区设置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName)
	// 与数据库建立连接,生成一个*gorm.DB类型的对象
	GORM, err = gorm.Open(config.DbType, dsn)
	if err != nil {
		logger.Error(errors.New("数据库连接失败,错误信息," + err.Error()))

		// 打印sql语句
		GORM.LogMode(config.LogMode)

		/* 开启连接池*/
		// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
		GORM.DB().SetMaxIdleConns(config.MaxIdleConns)
		// 设置了连接可复用的最大时间
		GORM.DB().SetMaxOpenConns(config.MaxOpenConns)
		// 设置了连接可复用的最大时间
		GORM.DB().SetConnMaxLifetime(time.Duration(config.MaxLifeTime))

		isInit = true
		logger.Info("连接数据库 " + config.DbHost + ":" + strconv.Itoa(config.DbPort) + "成功!")
	}
}

// Close 关闭数据库连接
func Close() error {
	return GORM.Close()
}
