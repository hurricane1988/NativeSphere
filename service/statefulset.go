package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/wonderivan/logger"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StatefulSet 实例化statefulSet结构体
var StatefulSet statefulSet

// 定义statefulSet结构体
type statefulSet struct{}

// StatefulSetsResp 定义statefulSet原数据
type StatefulSetsResp struct {
	Items []appsv1.StatefulSet `json:"items"`
	Total int                  `json:"total"`
}

// GetStatefulSets 获取statefulSets列表、支持过滤、排序、分页
func (s *statefulSet) GetStatefulSets(filterName, namespace string, limit, page int) (statefulSetsResp *StatefulSetsResp, err error) {
	// 获取statefulSetList类型的statefulSet
	statefulSetList, err := K8s.ClientSet.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(errors.New("获取statefulSet列表失败，错误信息" + err.Error()))
		return nil, errors.New("获取statefulSet列表失败，错误信息" + err.Error())
	}
	// 将statefulSetList中的StatefulSet列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataSelector{
		GenericDataList: s.toCells(statefulSetList.Items),
		DataSelectQuery: &DataSelectQuery{
			Filter: &FilterQuery{Name: filterName},
			Paginate: &PaginateQuery{
				limit,
				page,
			},
		},
	}
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	data := filtered.Sort().Paginate()

	// 将[]DataCell类型的statefulSet列表转为v1.statefulset列表
	statefulSets := s.fromCells(data.GenericDataList)

	return &StatefulSetsResp{
		Items: statefulSets,
		Total: total,
	}, nil
}

// GetStatefulSetDetail 获取statefulSets详情
func (s *statefulSet) GetStatefulSetDetail(statefulSetName, namespace string) (statefulSet *appsv1.StatefulSet, err error) {
	statefulSet, err = K8s.ClientSet.AppsV1().StatefulSets(namespace).Get(context.TODO(), statefulSetName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取statefulSet" + statefulSetName + "详情失败,错误信息 " + err.Error()))
		return nil, errors.New("获取statefulSet" + statefulSetName + "详情失败,错误信息 " + err.Error())
	}
	return statefulSet, nil
}

// DeleteStatefulSet 删除statefulSet
func (s *statefulSet) DeleteStatefulSet(statefulSetName, namespace string) (err error) {
	err = K8s.ClientSet.AppsV1().StatefulSets(namespace).Delete(context.TODO(), statefulSetName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除StatefulSet失败" + statefulSetName + "失败，错误信息" + err.Error()))
		return errors.New("删除StatefulSet失败" + statefulSetName + "失败，错误信息" + err.Error())
	}

	return nil
}

// UpdateStatefulSet 更新statefulSet
func (s *statefulSet) UpdateStatefulSet(namespace, content string) (err error) {
	var statefulSet = &appsv1.StatefulSet{}

	err = json.Unmarshal([]byte(content), statefulSet)
	if err != nil {
		logger.Error(errors.New("反序列化失败, " + err.Error()))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = K8s.ClientSet.AppsV1().StatefulSets(namespace).Update(context.TODO(), statefulSet, metav1.UpdateOptions{})
	if err != nil {
		logger.Error(errors.New("更新StatefulSet失败, " + err.Error()))
		return errors.New("更新StatefulSet失败, " + err.Error())
	}
	return nil
}

func (s *statefulSet) toCells(std []appsv1.StatefulSet) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = statefulSetCell(std[i])
	}
	return cells
}

func (s *statefulSet) fromCells(cells []DataCell) []appsv1.StatefulSet {
	statefulSets := make([]appsv1.StatefulSet, len(cells))
	for i := range cells {
		statefulSets[i] = appsv1.StatefulSet(cells[i].(statefulSetCell))
	}

	return statefulSets
}
