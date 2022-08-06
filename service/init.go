package service

import (
	"NativeSphere/config"
	"github.com/wonderivan/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 用于初始化k8s clientset

// K8s 实例化k8s
var K8s k8s

// 生产k8s结构体方法
type k8s struct {
	ClientSet *kubernetes.Clientset
}

// Init 初始化k8s
func (k *k8s) Init() {
	conf, err := clientcmd.BuildConfigFromFlags("", config.Kubeconfig)
	if err != nil {
		logger.Error("初始化k8s配置失败," + err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(conf)
	if err != nil {
		logger.Error("初始化k8s clientSet失败， " + err.Error())
	} else {
		logger.Info("初始化k8s clientSet成功!")
	}
	k.ClientSet = clientSet
}
