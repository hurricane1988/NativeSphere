package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigMap 定义configmap变量
var ConfigMap configMap

// 定义configmap结构体
type configMap struct{}

type ConfigMapsResp struct {
	Items []corev1.ConfigMap `json:"items"`
	Total int                `json:"total"`
}

// GetConfigMaps 获取configmap列表、支持过滤、排序、分页
func (c *configMap) GetConfigMaps(filterName, namespace string, limit, page int) (configMapsResp *ConfigMapsResp, err error) {
	// 获取configmapList类型的configMap
	configMapList, err := K8s.ClientSet.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(errors.New("获取namespace " + namespace + "下configMap " + filterName + "失败,错误信息 " + err.Error()))
		return nil, errors.New("获取namespace " + namespace + "下configMap " + filterName + "失败,错误信息 " + err.Error())
	}
	// 将configMapList中的configMap列表(items),放进dataselector对象中,进行排序
	selectableData := &dataSelector{
		GenericDataList: c.toCells(configMapList.Items),
		DataSelectQuery: &DataSelectQuery{
			Filter: &FilterQuery{
				Name: filterName,
			},
			Paginate: &PaginateQuery{
				limit, page,
			},
		},
	}
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	data := filtered.Sort().Paginate()

	// 将[]DataCell类型的configmap列表转为v1.configmap列表
	configMaps := c.fromCells(data.GenericDataList)

	return &ConfigMapsResp{
		Items: configMaps,
		Total: total,
	}, nil
}

// GetConfigMapDetail 获取configMap详情
func (c *configMap) GetConfigMapDetail(configMapName, namespace string) (configMap *corev1.ConfigMap, err error) {
	configMap, err = K8s.ClientSet.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取namespace下 " + namespace + "configMap " + configMapName + "失败,错误信息 " + err.Error()))
		return nil, errors.New("获取namespace下 " + namespace + "configMap " + configMapName + "失败,错误信息 " + err.Error())
	}
	return configMap, nil
}

// DeleteConfigMap 删除configMap
func (c *configMap) DeleteConfigMap(configMapName, namespace string) (err error) {
	err = K8s.ClientSet.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), configMapName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除ConfigMap " + configMapName + "失败,错误信息 " + err.Error()))
		return errors.New("删除ConfigMap " + configMapName + "失败,错误信息 " + err.Error())
	}
	return nil
}

// UpdateConfigMap 更新configmap
func (c *configMap) UpdateConfigMap(namespace, content string) (err error) {
	var configMap = &corev1.ConfigMap{}
	err = json.Unmarshal([]byte(content), configMap)
	if err != nil {
		logger.Error(errors.New("反序列化失败，错误信息 " + err.Error()))
		return errors.New("反序列化失败，错误信息 " + err.Error())
	}
	_, err = K8s.ClientSet.CoreV1().ConfigMaps(namespace).Update(context.TODO(), configMap, metav1.UpdateOptions{})
	if err != nil {
		logger.Error(errors.New("更新ConfigMap失败,错误信息 " + err.Error()))
		return errors.New("更新ConfigMap失败,错误信息 " + err.Error())
	}
	return nil
}

//
func (c *configMap) toCells(std []corev1.ConfigMap) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = configMapCell(std[i])
	}
	return cells
}

//
func (c *configMap) fromCells(cells []DataCell) []corev1.ConfigMap {
	configMaps := make([]corev1.ConfigMap, len(cells))
	for i := range cells {
		configMaps[i] = corev1.ConfigMap(cells[i].(configMapCell))
	}
	return configMaps
}
